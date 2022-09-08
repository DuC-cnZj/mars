package wssender

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/models"
	"google.golang.org/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/duc-cnzj/mars/internal/contracts"

	websocket_pb "github.com/duc-cnzj/mars-client/v4/websocket"

	"github.com/duc-cnzj/mars/internal/adapter"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/plugins"
	gonsq "github.com/nsqio/go-nsq"
)

const ephemeralBroadroom = BroadcastRoom + "#ephemeral"

var nsqSenderName = "ws_sender_nsq"

func init() {
	dr := &NsqSender{}
	plugins.RegisterPlugin(dr.Name(), dr)
}

func getNsqProjectEventRoom[T int64 | int](nsID T) string {
	return fmt.Sprintf("project-pod-events-%d#ephemeral", nsID)
}

type NsqSender struct {
	producer    *gonsq.Producer
	cfg         *gonsq.Config
	lookupdAddr string
	addr        string
}

func (n *NsqSender) Name() string {
	return nsqSenderName
}

func (n *NsqSender) Initialize(args map[string]any) (err error) {
	n.cfg = gonsq.NewConfig()
	if s, ok := args["addr"]; ok {
		n.addr = s.(string)
	} else {
		err = errors.New("[nsq]: add not exits")
		return
	}
	if s, ok := args["lookupdAddr"]; ok {
		n.lookupdAddr = s.(string)
	}
	p, err := gonsq.NewProducer(n.addr, n.cfg)
	if err != nil {
		return err
	}
	setLogLevel(p)
	err = p.Ping()
	n.producer = p
	mlog.Info("[Plugin]: " + n.Name() + " plugin Initialize...")
	return
}

func (n *NsqSender) Destroy() error {
	n.producer.Stop()
	mlog.Info("[Plugin]: " + n.Name() + " plugin Destroy...")
	return nil
}

func (n *NsqSender) New(uid, id string) contracts.PubSub {
	return &nsq{
		addr:        n.addr,
		lookupdAddr: n.lookupdAddr,
		cfg:         n.cfg,
		uid:         uid,
		id:          id,
		consumers:   map[string]*gonsq.Consumer{},
		producer:    n.producer,
		msgCh:       make(chan []byte, messageChSize),
		eventMsgCh:  make(chan []byte, messageChSize),
		channels:    map[string]struct{}{},
	}
}

type nsq struct {
	addr, lookupdAddr string
	cfg               *gonsq.Config
	uid, id           string

	consumersMu sync.RWMutex
	consumers   map[string]*gonsq.Consumer

	producer   *gonsq.Producer
	msgCh      chan []byte
	eventMsgCh chan []byte

	mu       sync.RWMutex
	channels map[string]struct{}
}

type jsonHandler struct {
	ch chan []byte
}

func (j *jsonHandler) HandleMessage(m *gonsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	j.ch <- m.Body

	return nil
}

func (n *nsq) Join(projectID int64) error {
	var pmodel models.Project
	if err := app.DB().First(&pmodel, projectID).Error; err != nil {
		return err
	}
	channel := getNsqProjectEventRoom(pmodel.NamespaceId)

	consumer, err := gonsq.NewConsumer(channel, n.ephemeralID(), n.cfg)
	if err != nil {
		mlog.Error(err)
		return err
	}
	if err := connect(consumer, n.addr, n.lookupdAddr, &jsonHandler{ch: n.eventMsgCh}); err != nil {
		mlog.Error(err)
		return err
	}

	func() {
		n.consumersMu.Lock()
		defer n.consumersMu.Unlock()
		n.consumers[channel] = consumer
	}()

	func() {
		n.mu.Lock()
		defer n.mu.Unlock()
		n.channels[channel] = struct{}{}
	}()
	return nil
}

func (n *nsq) Leave(nsID int64, projectID int64) error {
	channel := getNsqProjectEventRoom(nsID)

	func() {
		n.consumersMu.Lock()
		defer n.consumersMu.Unlock()
		consumer, ok := n.consumers[channel]
		if ok {
			consumer.Stop()
			delete(n.consumers, channel)
		}
	}()

	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.channels, channel)
	return nil
}

func (n *nsq) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-n.eventMsgCh:
			if !ok {
				mlog.Warning("nsq event ch closed")
				return nil
			}
			var obj projectEventObj
			if err := json.Unmarshal([]byte(data), &obj); err != nil {
				mlog.Error(err)
			}
			fn := func() bool {
				n.mu.RLock()
				defer n.mu.RUnlock()
				_, ok := n.channels[obj.Channel]
				return ok
			}
			if !fn() {
				continue
			}

			var projects []models.Project
			if err := app.DB().Select("id", "namespace_id", "pod_selectors").Where("`namespace_id` = ?", obj.NamespaceID).Find(&projects).Error; err != nil {
				mlog.Error(err)
			}
			for _, project := range projects {
				func() {
					for _, s := range project.GetPodSelectors() {
						parse, _ := labels.Parse(s)
						if parse.Matches(labels.Set(obj.Pod.Labels)) {
							marshal, _ := proto.Marshal(&websocket_pb.WsProjectPodEventResponse{
								Metadata: &websocket_pb.Metadata{
									Type:   websocket_pb.Type_ProjectPodEvent,
									End:    true,
									Result: websocket_pb.ResultType_Success,
									To:     plugins.ToSelf,
								},
								ProjectId: int64(project.ID),
							})
							n.msgCh <- marshal
							return
						}
					}
				}()
			}
		}
	}
}

func (n *nsq) Publish(nsID int64, pod *v1.Pod) error {
	room := getNsqProjectEventRoom(nsID)
	marshal, err := json.Marshal(&projectEventObj{
		NamespaceID: nsID,
		Pod:         pod,
		Channel:     room,
	})
	if err != nil {
		mlog.Error(err)
	}
	return n.producer.Publish(room, marshal)
}

func (n *nsq) Info() any {
	return nil
}

func (n *nsq) Uid() string {
	return n.uid
}

func (n *nsq) ID() string {
	return n.id
}

func (n *nsq) ToSelf(response contracts.WebsocketMessage) error {
	return n.producer.Publish(n.ephemeralID(), plugins.ProtoToMessage(response, websocket_pb.To_ToSelf, n.id).Marshal())
}

func (n *nsq) ToAll(response contracts.WebsocketMessage) error {
	return n.producer.Publish(ephemeralBroadroom, plugins.ProtoToMessage(response, websocket_pb.To_ToAll, n.id).Marshal())
}

func (n *nsq) ToOthers(response contracts.WebsocketMessage) error {
	return n.producer.Publish(ephemeralBroadroom, plugins.ProtoToMessage(response, websocket_pb.To_ToOthers, n.id).Marshal())
}

func (n *nsq) ephemeralID() string {
	return n.ID() + "#ephemeral"
}

func (n *nsq) Subscribe() <-chan []byte {
	consumerAll, _ := gonsq.NewConsumer(ephemeralBroadroom, n.ephemeralID(), n.cfg)
	consumer, _ := gonsq.NewConsumer(n.ephemeralID(), n.ephemeralID(), n.cfg)
	h := &handler{msgCh: n.msgCh, id: n.id}
	connect(consumer, n.addr, n.lookupdAddr, h)
	connect(consumerAll, n.addr, n.lookupdAddr, h)

	n.consumersMu.Lock()
	defer n.consumersMu.Unlock()
	n.consumers = map[string]*gonsq.Consumer{
		ephemeralBroadroom: consumerAll,
		n.ephemeralID():    consumer,
	}

	return n.msgCh
}

func connect(consumer *gonsq.Consumer, addr, lookupdAddr string, h gonsq.Handler) error {
	setLogLevel(consumer)
	consumer.AddHandler(h)

	var err error
	if lookupdAddr != "" {
		err = consumer.ConnectToNSQLookupd(lookupdAddr)
	} else {
		err = consumer.ConnectToNSQD(addr)
	}
	return err
}

func (n *nsq) Close() error {
	defer mlog.Debugf("[nsq]: id: %v closed", n.ID())
	for _, c := range n.consumers {
		c.Stop()
		if n.lookupdAddr != "" {
			c.DisconnectFromNSQLookupd(n.lookupdAddr)
		} else {
			c.DisconnectFromNSQD(n.addr)
		}
	}
	return nil
}

type handler struct {
	id    string
	msgCh chan []byte
}

func (h *handler) HandleMessage(m *gonsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	message, _ := plugins.DecodeMessage(m.Body)
	switch message.To {
	case plugins.ToSelf:
		fallthrough
	case plugins.ToAll:
		h.msgCh <- message.Data
	case plugins.ToOthers:
		if message.ID != h.id {
			h.msgCh <- message.Data
		}
	}

	return nil
}

func setLogLevel(s any) {
	if ss, ok := s.(*gonsq.Consumer); ok {
		ss.SetLoggerLevel(gonsq.LogLevelError)
		ss.SetLoggerForLevel(&adapter.NsqLoggerAdapter{}, gonsq.LogLevelError)
	}
	if ss, ok := s.(*gonsq.Producer); ok {
		ss.SetLoggerLevel(gonsq.LogLevelError)
		ss.SetLoggerForLevel(&adapter.NsqLoggerAdapter{}, gonsq.LogLevelError)
	}
}

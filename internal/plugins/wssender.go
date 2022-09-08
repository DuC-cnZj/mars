package plugins

//go:generate mockgen -destination ../mock/mock_wssender.go -package mock github.com/duc-cnzj/mars/internal/plugins WsSender

import (
	"context"
	"encoding/json"
	"sync"

	v1 "k8s.io/api/core/v1"

	"google.golang.org/protobuf/proto"

	websocket_pb "github.com/duc-cnzj/mars-client/v4/websocket"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
)

var wsSenderOnce sync.Once

const (
	ToSelf   = websocket_pb.To_ToSelf
	ToAll    = websocket_pb.To_ToAll
	ToOthers = websocket_pb.To_ToOthers
)

type Message struct {
	Data []byte
	To   websocket_pb.To
	ID   string
}

func (m Message) Marshal() []byte {
	marshal, _ := json.Marshal(&m)
	return marshal
}

func DecodeMessage(data []byte) (msg Message, err error) {
	err = json.Unmarshal(data, &msg)
	return
}

func ProtoToMessage(m proto.Message, to websocket_pb.To, id string) Message {
	marshal, _ := proto.Marshal(m)

	return Message{
		Data: marshal,
		To:   to,
		ID:   id,
	}
}

type WsMetadataResponse = websocket_pb.WsMetadataResponse

var _ contracts.WebsocketMessage = (*WsMetadataResponse)(nil)

type WsSender interface {
	contracts.PluginInterface

	New(uid, id string) contracts.PubSub
}

func GetWsSender() WsSender {
	pcfg := app.Config().WsSenderPlugin
	p := app.App().GetPluginByName(pcfg.Name)
	wsSenderOnce.Do(func() {
		if err := p.Initialize(pcfg.GetArgs()); err != nil {
			panic(err)
		}
		app.App().RegisterAfterShutdownFunc(func(applicationInterface contracts.ApplicationInterface) {
			p.Destroy()
		})
	})

	return p.(WsSender)
}

type EmptyPubSub struct{}

func (e *EmptyPubSub) Join(projectID int64) error {
	return nil
}

func (e *EmptyPubSub) Leave(nsID int64, projectID int64) error {
	return nil
}

func (e *EmptyPubSub) Run(ctx context.Context) error {
	return nil
}

func (e *EmptyPubSub) Publish(nsID int64, pod *v1.Pod) error {
	//TODO implement me
	panic("implement me")
}

func (e *EmptyPubSub) Info() any {
	return nil
}

func (e *EmptyPubSub) Uid() string {
	return ""
}

func (e *EmptyPubSub) ID() string {
	return ""
}

func (e *EmptyPubSub) ToSelf(message contracts.WebsocketMessage) error {
	return nil
}

func (e *EmptyPubSub) ToAll(message contracts.WebsocketMessage) error {
	return nil
}

func (e *EmptyPubSub) ToOthers(message contracts.WebsocketMessage) error {
	return nil
}

func (e *EmptyPubSub) Subscribe() <-chan []byte {
	return nil
}

func (e *EmptyPubSub) Close() error {
	return nil
}

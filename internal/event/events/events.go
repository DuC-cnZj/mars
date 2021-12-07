package events

import (
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/plugins"
	websocket_pb "github.com/duc-cnzj/mars/pkg/websocket"

	v1 "k8s.io/api/core/v1"
)

var (
	EventNamespaceCreated contracts.Event = "namespace_created"
	EventNamespaceDeleted contracts.Event = "namespace_deleted"

	EventProjectedDeleted contracts.Event = "projected_deleted"
)

type NamespaceCreatedData struct {
	NsModel  *models.Namespace
	NsK8sObj *v1.Namespace
}

type NamespaceDeletedData struct {
	NsModel *models.Namespace
}

func HandleNamespaceDeleted(data interface{}, e contracts.Event) error {
	plugins.GetWsSender().New("", "").ToAll(&plugins.WsResponseMetadata{Metadata: &websocket_pb.ResponseMetadata{Type: websocket_pb.Type_ReloadProjects}})
	mlog.Debug("event handled: ", e.String())

	return nil
}

func HandleProjectDeleted(data interface{}, e contracts.Event) error {
	plugins.GetWsSender().New("", "").ToAll(&plugins.WsResponseMetadata{Metadata: &websocket_pb.ResponseMetadata{Type: websocket_pb.Type_ReloadProjects}})
	mlog.Debug("event handled: ", e.String(), data)

	return nil
}

package socket

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes/fake"
	fake2 "k8s.io/metrics/pkg/client/clientset/versioned/fake"

	"github.com/duc-cnzj/mars-client/v4/types"
	"github.com/duc-cnzj/mars-client/v4/websocket"
	auth2 "github.com/duc-cnzj/mars/internal/auth"
	"github.com/duc-cnzj/mars/internal/cache_lock"
	"github.com/duc-cnzj/mars/internal/config"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/mock"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/testutil"
	"github.com/duc-cnzj/mars/internal/utils"
)

func TestHandleWsAuthorize(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	privateKey, _ := x509.MarshalPKCS8PrivateKey(key)
	bf := bytes.Buffer{}
	pem.Encode(&bf, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey})
	authSvc := auth2.NewJwtAuth(key, key.Public().(*rsa.PublicKey))
	app := testutil.MockApp(m)
	app.EXPECT().Auth().Return(authSvc).AnyTimes()
	sign, _ := authSvc.Sign(contracts.UserInfo{
		Name: "duc",
	})
	marshal, _ := proto.Marshal(&websocket.AuthorizeTokenInput{Token: sign.Token})
	pubsub := mock.NewMockPubSub(m)
	pubsub.EXPECT().ToSelf(gomock.Any()).Times(1)
	conn := &WsConn{
		pubSub: pubsub,
	}
	HandleWsAuthorize(conn, websocket.Type_HandleAuthorize, []byte("1:"))
	HandleWsAuthorize(conn, websocket.Type_HandleAuthorize, marshal)
	assert.Equal(t, "duc", conn.GetUser().Name)
}

func TestHandleWsCancel(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()

	pubsub := mock.NewMockPubSub(m)
	pubsub.EXPECT().ToSelf(gomock.Any()).Times(1)

	app := testutil.MockApp(m)
	db, fn := testutil.SetGormDB(m, app)
	defer fn()
	ns := &models.Namespace{
		Name: "ns",
	}
	db.Create(ns)
	marshal, _ := proto.Marshal(&websocket.CancelInput{
		Type:        websocket.Type_CancelProject,
		NamespaceId: int64(ns.ID),
		Name:        "app",
	})
	cs := mock.NewMockCancelSignaler(m)
	slug := utils.GetSlugName(ns.ID, "app")
	cs.EXPECT().Has(slug).Return(true).Times(1)
	cs.EXPECT().Cancel(slug).Times(1)
	testutil.AssertAuditLogFired(m, app)
	conn := &WsConn{pubSub: pubsub, cancelSignaler: cs}
	HandleWsCancel(conn, websocket.Type_CancelProject, []byte("1:"))
	HandleWsCancel(conn, websocket.Type_CancelProject, marshal)
}

func TestHandleWsCreateProject(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	cs := mock.NewMockCancelSignaler(m)
	sm := mock.NewMockSessionMapper(m)
	ps := mock.NewMockPubSub(m)
	app := testutil.MockApp(m)
	l := mock.NewMockLocker(m)
	app.EXPECT().CacheLock().Return(l).AnyTimes()
	l.EXPECT().RenewalAcquire(gomock.Any(), int64(30), int64(20)).Return(func() {}, true).AnyTimes()
	c := &WsConn{
		id:  "id",
		uid: "uid",
		NewJobFunc: func(input *JobInput, user contracts.UserInfo, slugName string, messager contracts.DeployMsger, pubsub contracts.PubSub, timeoutSeconds int64, opts ...Option) contracts.Job {
			return job
		},
		cancelSignaler:   cs,
		terminalSessions: sm,
		pubSub:           ps,
	}

	msg := mock.NewMockDeployMsger(m)

	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsCreateProject(c, websocket.Type_UpdateProject, []byte("1:"))

	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().IsDryRun().Return(false).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(nil).Times(1)
	job.EXPECT().Run().Return(nil).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().ID().Return("1").AnyTimes()
	cs.EXPECT().Add("1", gomock.Any()).Return(nil)
	job.EXPECT().AddDestroyFunc(gomock.Any()).AnyTimes()

	marshal, _ := proto.Marshal(&websocket.CreateProjectInput{
		Type:         websocket.Type_CreateProject,
		NamespaceId:  0,
		Name:         "",
		GitProjectId: 0,
		GitBranch:    "",
		GitCommit:    "",
		Config:       "",
		Atomic:       false,
		ExtraValues:  nil,
	})

	HandleWsCreateProject(c, websocket.Type_CreateProject, marshal)

	cs.EXPECT().Add("1", gomock.Any()).Return(errors.New("xxx"))
	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsCreateProject(c, websocket.Type_CreateProject, marshal)
}

func TestHandleWsHandleCloseShell(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	sm := mock.NewMockSessionMapper(m)
	pubsub := mock.NewMockPubSub(m)
	pubsub.EXPECT().ToSelf(gomock.Any()).Times(1)
	conn := &WsConn{
		pubSub:           pubsub,
		terminalSessions: sm,
	}
	marshal, _ := proto.Marshal(&websocket.TerminalMessageInput{
		Type: websocket.Type_HandleCloseShell,
		Message: &websocket.TerminalMessage{
			SessionId: "sid",
		},
	})
	sm.EXPECT().Close("sid", uint32(0), "").Times(1)
	HandleWsHandleCloseShell(conn, websocket.Type_HandleCloseShell, []byte("1234:"))
	HandleWsHandleCloseShell(conn, websocket.Type_HandleCloseShell, marshal)
}

func TestHandleWsHandleExecShell(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	pubsub := mock.NewMockPubSub(m)
	pubsub.EXPECT().ToSelf(&websocket.WsHandleShellResponse{
		Metadata: &websocket.Metadata{
			Id:     "id-x",
			Uid:    "uid-x",
			Type:   WsHandleExecShell,
			Result: ResultSuccess,
		},
		TerminalMessage: &websocket.TerminalMessage{
			SessionId: "id",
		},
		Container: &types.Container{
			Namespace: "ns",
			Pod:       "pod",
			Container: "app",
		},
	}).Times(1)
	conn := &WsConn{
		id:     "id-x",
		uid:    "uid-x",
		pubSub: pubsub,
		NewShellFunc: func(input *websocket.WsHandleExecShellInput, conn *WsConn) (string, error) {
			return "id", nil
		},
	}
	marshal, _ := proto.Marshal(&websocket.WsHandleExecShellInput{
		Type: websocket.Type_HandleExecShell,
		Container: &types.Container{
			Namespace: "ns",
			Pod:       "pod",
			Container: "app",
		},
	})
	HandleWsHandleExecShell(conn, websocket.Type_HandleExecShell, marshal)
	conn.NewShellFunc = func(input *websocket.WsHandleExecShellInput, conn *WsConn) (string, error) {
		return "", errors.New("xxx")
	}
	pubsub.EXPECT().ToSelf(&websocket.WsMetadataResponse{
		Metadata: &websocket.Metadata{
			Slug:    "",
			Type:    websocket.Type_HandleExecShell,
			Result:  ResultError,
			End:     true,
			Uid:     "uid-x",
			Id:      "id-x",
			Message: "xxx",
		},
	}).Times(1)
	HandleWsHandleExecShell(conn, websocket.Type_HandleExecShell, marshal)
	pubsub.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsHandleExecShell(conn, websocket.Type_HandleExecShell, []byte("1234:"))
}

type protoMatcher struct {
	wants proto.Message
}

func (m *protoMatcher) Matches(x any) bool {
	return proto.Equal(x.(proto.Message), m.wants)
}

func (m *protoMatcher) String() string {
	return ""
}

func TestHandleWsHandleExecShellMsg(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	sm := mock.NewMockSessionMapper(m)
	ps := mock.NewMockPubSub(m)
	conn := &WsConn{
		terminalSessions: sm,
		pubSub:           ps,
	}
	message := &websocket.TerminalMessage{
		Data:      []byte("data"),
		SessionId: "sid",
	}
	marshal, _ := proto.Marshal(&websocket.TerminalMessageInput{
		Type:    websocket.Type_HandleExecShellMsg,
		Message: message,
	})
	sm.EXPECT().Send(&protoMatcher{wants: proto.Clone(message)}).Times(1)
	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsHandleExecShellMsg(conn, websocket.Type_HandleExecShellMsg, []byte("1:"))
	HandleWsHandleExecShellMsg(conn, websocket.Type_HandleExecShellMsg, marshal)
}

func TestHandleWsUpdateProjectSignalError(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	cs := &CancelSignals{
		cs: map[string]func(error){},
	}
	ps := mock.NewMockPubSub(m)
	c := &WsConn{
		id:  "id",
		uid: "uid",
		NewJobFunc: func(input *JobInput, user contracts.UserInfo, slugName string, messager contracts.DeployMsger, pubsub contracts.PubSub, timeoutSeconds int64, opts ...Option) contracts.Job {
			return job
		},
		cancelSignaler: cs,
		pubSub:         ps,
	}

	app := testutil.MockApp(m)
	db, f := testutil.SetGormDB(m, app)
	defer f()
	db.AutoMigrate(&models.Project{}, &models.Namespace{})
	p := &models.Project{Name: "app", Namespace: models.Namespace{Name: "ns"}}
	assert.Nil(t, db.Create(p).Error)
	marshal, _ := proto.Marshal(&websocket.UpdateProjectInput{
		Type:      websocket.Type_UpdateProject,
		ProjectId: int64(p.ID),
	})
	slug := utils.GetSlugName(p.Namespace.ID, p.Name)
	ps.EXPECT().ToSelf(&websocket.WsMetadataResponse{
		Metadata: &websocket.Metadata{
			Slug:    slug,
			Type:    websocket.Type_UpdateProject,
			Result:  ResultDeployFailed,
			End:     true,
			Uid:     "uid",
			Id:      "id",
			Message: "正在清理中，请稍后再试。",
		},
	}).Times(1)
	job.EXPECT().ID().Return(slug).AnyTimes()
	job.EXPECT().IsDryRun().Return(false).AnyTimes()
	cs.Add(slug, func(err error) {})
	job.EXPECT().AddDestroyFunc(gomock.Any()).AnyTimes()

	l := mock.NewMockLocker(m)
	l.EXPECT().RenewalAcquire(gomock.Any(), int64(30), int64(20)).Return(func() {}, true).AnyTimes()
	app.EXPECT().CacheLock().Return(l).AnyTimes()
	HandleWsUpdateProject(c, websocket.Type_UpdateProject, marshal)
}

func TestHandleWsUpdateProject(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	cs := mock.NewMockCancelSignaler(m)
	sm := mock.NewMockSessionMapper(m)
	ps := mock.NewMockPubSub(m)
	app := testutil.MockApp(m)
	l := mock.NewMockLocker(m)
	app.EXPECT().CacheLock().Return(l).AnyTimes()
	l.EXPECT().RenewalAcquire(gomock.Any(), int64(30), int64(20)).Return(func() {}, true).AnyTimes()
	//job.EXPECT().AddDestroyFunc(gomock.Any()).AnyTimes()
	job.EXPECT().IsDryRun().Return(false).AnyTimes()
	c := &WsConn{
		id:  "id",
		uid: "uid",
		NewJobFunc: func(input *JobInput, user contracts.UserInfo, slugName string, messager contracts.DeployMsger, pubsub contracts.PubSub, timeoutSeconds int64, opts ...Option) contracts.Job {
			return job
		},
		cancelSignaler:   cs,
		terminalSessions: sm,
		pubSub:           ps,
	}

	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsUpdateProject(c, websocket.Type_UpdateProject, []byte("1:"))

	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(nil).Times(1)
	job.EXPECT().Run().Return(nil).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().ID().Return("1").AnyTimes()
	cs.EXPECT().Add("1", gomock.Any()).Return(nil)

	manager := mock.NewMockDBManager(m)
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.Exec("PRAGMA foreign_keys = ON", nil)
	s, _ := db.DB()
	defer s.Close()
	manager.EXPECT().DB().Return(db).AnyTimes()
	app.EXPECT().DBManager().Return(manager).AnyTimes()
	db.AutoMigrate(&models.Project{}, &models.Namespace{})

	targetProject := &models.Project{
		Name:      "app",
		Namespace: models.Namespace{Name: "ns"},
	}
	db.Create(targetProject)

	marshal, _ := proto.Marshal(&websocket.UpdateProjectInput{
		Type:        websocket.Type_UpdateProject,
		ProjectId:   int64(targetProject.ID),
		GitBranch:   "",
		GitCommit:   "",
		Config:      "",
		Atomic:      false,
		ExtraValues: nil,
	})

	marshalNotFound, _ := proto.Marshal(&websocket.UpdateProjectInput{
		Type:        websocket.Type_UpdateProject,
		ProjectId:   99999,
		GitBranch:   "",
		GitCommit:   "",
		Config:      "",
		Atomic:      false,
		ExtraValues: nil,
	})
	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	job.EXPECT().AddDestroyFunc(gomock.Any()).AnyTimes()
	HandleWsUpdateProject(c, websocket.Type_UpdateProject, marshalNotFound)

	HandleWsUpdateProject(c, websocket.Type_UpdateProject, marshal)

	cs.EXPECT().Add(gomock.Any(), gomock.Any()).Return(errors.New(""))
	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsUpdateProject(c, websocket.Type_UpdateProject, marshal)
}

func TestInstallProject_Success(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(nil).Times(1)
	job.EXPECT().Run().Return(nil).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	InstallProject(job)
}

func TestInstallProject_ValidateFail(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	msg.EXPECT().SendEndError(gomock.Any()).Times(1)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(errors.New("err Validate")).Times(1)
	job.EXPECT().GetStoppedErrorIfHas().Return(nil).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().IsDryRun().Return(false).Times(1)
	job.EXPECT().Prune().Times(1)
	assert.Equal(t, "err Validate", InstallProject(job).Error())
}
func TestInstallProject_ValidateFail2(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().SetDeployResult(websocket.ResultType_DeployedCanceled, "stopped", nil)
	job.EXPECT().ProjectModel().Return(nil).Times(1)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(errors.New("err Validate")).Times(1)
	job.EXPECT().GetStoppedErrorIfHas().Return(errors.New("stopped")).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().IsDryRun().Return(false).Times(1)
	job.EXPECT().Prune().Times(1)
	assert.Equal(t, "stopped", InstallProject(job).Error())
}

func TestInstallProject_RunFail(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(nil).Times(1)
	job.EXPECT().Run().Return(errors.New("err")).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().IsDryRun().Return(false).Times(1)
	job.EXPECT().Prune().Times(1)
	InstallProject(job)
}

func TestInstallProject_LoadConfigsFail(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(errors.New("")).Times(1)
	msg.EXPECT().SendEndError(gomock.Any()).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().IsDryRun().Return(false).Times(1)
	job.EXPECT().Prune().Times(1)
	job.EXPECT().GetStoppedErrorIfHas().Return(nil).Times(1)
	InstallProject(job)
}

func TestInstallProject_LoadConfigsFail2(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	job := mock.NewMockJob(m)
	msg := mock.NewMockDeployMsger(m)
	job.EXPECT().Messager().Return(msg).AnyTimes()
	job.EXPECT().Validate().Return(nil).Times(1)
	job.EXPECT().LoadConfigs().Return(errors.New("")).Times(1)
	job.EXPECT().CallDestroyFuncs().Times(1)
	job.EXPECT().Finish().Times(1)
	job.EXPECT().IsDryRun().Return(false).Times(1)
	job.EXPECT().Prune().Times(1)
	job.EXPECT().GetStoppedErrorIfHas().Return(errors.New("err")).Times(1)
	job.EXPECT().SetDeployResult(ResultDeployCanceled, "err", gomock.Any()).Times(1)
	job.EXPECT().ProjectModel().Return(nil)
	InstallProject(job)
}

func TestNewWebsocketManager(t *testing.T) {
	assert.NotNil(t, NewWebsocketManager(1*time.Second))
}

type rw struct {
	h http.Header
	w []byte
	http.ResponseWriter
}

func (receiver *rw) Header() http.Header {
	return receiver.h
}

func (receiver *rw) Write(in []byte) (int, error) {
	receiver.w = in
	return len(in), nil
}

func TestWebsocketManager_Info(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	app := testutil.MockApp(m)
	app.EXPECT().Config().Return(&config.Config{
		WsSenderPlugin: config.Plugin{
			Name: "test_ws",
		},
	})
	ws := mock.NewMockWsSender(m)
	ps := mock.NewMockPubSub(m)
	ws.EXPECT().New("", "").Return(ps)
	ws.EXPECT().Initialize(gomock.Any()).AnyTimes()
	app.EXPECT().RegisterAfterShutdownFunc(gomock.Any()).AnyTimes()
	ps.EXPECT().Info().Return("info...")
	app.EXPECT().GetPluginByName("test_ws").Return(ws)
	rwer := &rw{
		h: map[string][]string{},
		w: nil,
	}
	NewWebsocketManager(1*time.Second).Info(rwer, nil)
	assert.Equal(t, "application/json", rwer.h["Content-Type"][0])
	marshal, _ := json.Marshal("info...")
	assert.Equal(t, marshal, rwer.w)
}

type testWait struct {
	c int
}

func (t *testWait) Inc() {
	t.c++
}

func (t *testWait) Dec() {
	t.c--
}

func (t *testWait) Wait() {

}
func (t *testWait) Count() int {
	return t.c
}

func TestWebsocketManager_initConn(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	app := testutil.MockApp(m)
	app.EXPECT().Config().Return(&config.Config{
		WsSenderPlugin: config.Plugin{
			Name: "test_ws",
		},
	})
	ps := mock.NewMockPubSub(m)

	ws := mock.NewMockWsSender(m)
	ws.EXPECT().New("xxx", gomock.Any()).Return(ps)
	ws.EXPECT().Initialize(gomock.Any()).AnyTimes()
	app.EXPECT().RegisterAfterShutdownFunc(gomock.Any()).AnyTimes()
	app.EXPECT().GetPluginByName("test_ws").Return(ws)

	parse, _ := url.Parse("https://mars.local/ws?uid=xxx")
	r := &http.Request{
		URL: parse,
	}
	Wait = &testWait{}
	defer func() {
		Wait = NewWaitSocketExit()
	}()
	conn := NewWebsocketManager(1*time.Second).initConn(r, nil)
	assert.Equal(t, ps, conn.pubSub)
	assert.NotEmpty(t, conn.id)
	assert.Equal(t, "xxx", conn.uid)
	assert.Nil(t, conn.conn)
	assert.NotNil(t, conn.terminalSessions)
	assert.Equal(t, 1, Wait.Count())
}

func TestWebsocketManager_initConn2(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	app := testutil.MockApp(m)
	app.EXPECT().Config().Return(&config.Config{
		WsSenderPlugin: config.Plugin{
			Name: "test_ws",
		},
	})
	ps := mock.NewMockPubSub(m)

	ws := mock.NewMockWsSender(m)
	ws.EXPECT().New(gomock.Any(), gomock.Any()).Return(ps)
	ws.EXPECT().Initialize(gomock.Any()).AnyTimes()
	app.EXPECT().RegisterAfterShutdownFunc(gomock.Any()).AnyTimes()
	app.EXPECT().GetPluginByName("test_ws").Return(ws)

	parse, _ := url.Parse("https://mars.local/ws")
	r := &http.Request{
		URL: parse,
	}
	Wait = &testWait{}
	defer func() {
		Wait = NewWaitSocketExit()
	}()
	conn := NewWebsocketManager(1*time.Second).initConn(r, nil)
	assert.NotEmpty(t, conn.uid)
}

func TestWsConn_GetUser(t *testing.T) {
	c := &WsConn{}
	assert.IsType(t, contracts.UserInfo{}, c.GetUser())
}

func TestWsConn_SetUser(t *testing.T) {
	c := &WsConn{}
	u := contracts.UserInfo{
		Name: "aaa",
	}
	c.SetUser(u)
	assert.Equal(t, "aaa", c.user.Name)
}

func TestWsConn_Shutdown(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	cs := mock.NewMockCancelSignaler(m)
	cs.EXPECT().CancelAll().Times(1)
	tm := mock.NewMockSessionMapper(m)
	tm.EXPECT().CloseAll().Times(1)
	ps := mock.NewMockPubSub(m)
	ps.EXPECT().Close().Times(1)
	conn := mock.NewMockWebsocketConn(m)
	conn.EXPECT().Close().Times(1)
	var doneCalled bool
	c := &WsConn{
		doneFunc: func() {
			doneCalled = true
		},
		conn:             conn,
		cancelSignaler:   cs,
		pubSub:           ps,
		terminalSessions: tm,
	}
	Wait = &testWait{}
	defer func() {
		Wait = NewWaitSocketExit()
	}()
	c.Shutdown()
	assert.True(t, doneCalled)
	assert.Equal(t, -1, Wait.Count())
}

func TestWebsocketManager_TickClusterHealth(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	app := testutil.MockApp(m)
	ch := make(chan struct{})
	l := cache_lock.NewMemoryLock([2]int{-1, 100}, cache_lock.NewMemStore())
	app.EXPECT().CacheLock().Return(l).AnyTimes()
	app.EXPECT().Done().Return(ch).Times(1)
	go func() {
		time.Sleep(1500 * time.Millisecond)
		close(ch)
	}()

	app.EXPECT().Config().Return(&config.Config{
		WsSenderPlugin: config.Plugin{
			Name: "test_ws",
		},
	})
	app.EXPECT().K8sClient().Return(&contracts.K8sClient{Client: fake.NewSimpleClientset(), MetricsClient: fake2.NewSimpleClientset()}).AnyTimes()
	ps := mock.NewMockPubSub(m)
	ps.EXPECT().ToAll(gomock.Any()).Times(1)

	ws := mock.NewMockWsSender(m)
	ws.EXPECT().New("", "").Return(ps)
	ws.EXPECT().Initialize(gomock.Any()).AnyTimes()
	app.EXPECT().RegisterAfterShutdownFunc(gomock.Any()).AnyTimes()
	app.EXPECT().GetPluginByName("test_ws").Return(ws)

	NewWebsocketManager(1 * time.Second).TickClusterHealth()
	time.Sleep(2 * time.Second)
}

type slowLocker struct {
	contracts.Locker
}

func (s *slowLocker) Acquire(key string, seconds int64) bool {
	if s.Locker.Acquire(key, seconds) {
		time.Sleep(200 * time.Millisecond)
		return true
	}
	return false
}

func TestWebsocketManager_TickClusterHealth_Parallel(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	app := testutil.MockApp(m)
	ch := make(chan struct{})
	l := cache_lock.NewMemStore()
	app.EXPECT().CacheLock().Return(&slowLocker{Locker: cache_lock.NewMemoryLock([2]int{-1, 100}, l)}).AnyTimes()
	app.EXPECT().Done().Return(ch).AnyTimes()
	go func() {
		time.Sleep(1500 * time.Millisecond)
		close(ch)
	}()

	app.EXPECT().Config().Return(&config.Config{
		WsSenderPlugin: config.Plugin{
			Name: "test_ws",
		},
	}).AnyTimes()
	app.EXPECT().K8sClient().Return(&contracts.K8sClient{Client: fake.NewSimpleClientset(), MetricsClient: fake2.NewSimpleClientset()}).AnyTimes()
	ps := mock.NewMockPubSub(m)
	ps.EXPECT().ToAll(gomock.Any()).Times(1)

	ws := mock.NewMockWsSender(m)
	ws.EXPECT().New("", "").Return(ps).AnyTimes()
	ws.EXPECT().Initialize(gomock.Any()).AnyTimes()
	app.EXPECT().RegisterAfterShutdownFunc(gomock.Any()).AnyTimes()
	app.EXPECT().GetPluginByName("test_ws").Return(ws).AnyTimes()

	for i := 0; i < 10; i++ {
		go NewWebsocketManager(1 * time.Second).TickClusterHealth()
	}
	time.Sleep(2 * time.Second)
}

func Test_Upgrader(t *testing.T) {
	assert.True(t, upgrader.CheckOrigin(nil))
}

func TestHandleWsProjectPodEvent(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	ps := mock.NewMockPubSub(m)
	c := &WsConn{
		id:     "id",
		uid:    "uid",
		pubSub: ps,
	}
	marshal, _ := proto.Marshal(&websocket.ProjectPodEventJoinInput{
		Type:        websocket.Type_ProjectPodEvent,
		Join:        false,
		ProjectId:   1,
		NamespaceId: 1,
	})
	ps.EXPECT().Leave(int64(1), int64(1)).Times(1)
	HandleWsProjectPodEvent(c, websocket.Type_ProjectPodEvent, marshal)
	marshal2, _ := proto.Marshal(&websocket.ProjectPodEventJoinInput{
		Type:        websocket.Type_ProjectPodEvent,
		Join:        true,
		ProjectId:   2,
		NamespaceId: 2,
	})
	ps.EXPECT().Join(int64(2)).Times(1)
	HandleWsProjectPodEvent(c, websocket.Type_ProjectPodEvent, marshal2)

	ps.EXPECT().ToSelf(gomock.Any()).Times(1)
	HandleWsProjectPodEvent(c, websocket.Type_ProjectPodEvent, []byte("xxxxx"))
}

func TestWebsocketManager_Ws(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/test-ws", nil)
	NewWebsocketManager(1*time.Second).Ws(recorder, request)
	assert.Equal(t, 400, recorder.Code)
}

func Test_handleWsRead(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	ps := mock.NewMockPubSub(m)

	ps.EXPECT().ToSelf(&websocket.WsMetadataResponse{
		Metadata: &websocket.Metadata{
			Type:    websocket.Type_HandleAuthorize,
			Result:  ResultSuccess,
			End:     false,
			Uid:     "uid",
			Id:      "",
			Message: "认证中，请稍等~",
		},
	}).Times(1)
	handleWsRead(&WsConn{
		uid:    "uid",
		pubSub: ps,
		user: contracts.UserInfo{
			ID: "",
		},
	}, &websocket.WsRequestMetadata{
		Type: websocket.Type_ApplyProject,
	}, nil, map[websocket.Type]HandleRequestFunc{
		websocket.Type_ApplyProject: func(c *WsConn, t websocket.Type, message []byte) {
		},
	})

	called := false
	handleWsRead(&WsConn{
		id:     "id",
		uid:    "uid",
		pubSub: ps,
		user: contracts.UserInfo{
			ID: "id",
		},
	}, &websocket.WsRequestMetadata{
		Type: websocket.Type_ApplyProject,
	}, nil, map[websocket.Type]HandleRequestFunc{
		websocket.Type_ApplyProject: func(c *WsConn, t websocket.Type, message []byte) {
			called = true
		},
	})
	assert.True(t, called)

	app := testutil.MockApp(m)
	app.EXPECT().IsDebug().Return(true).Times(1)
	assert.PanicsWithValue(t, "errxxx", func() {
		handleWsRead(&WsConn{
			id:     "id",
			uid:    "uid",
			pubSub: ps,
			user: contracts.UserInfo{
				ID: "id",
			},
		}, &websocket.WsRequestMetadata{
			Type: websocket.Type_ApplyProject,
		}, nil, map[websocket.Type]HandleRequestFunc{
			websocket.Type_ApplyProject: func(c *WsConn, t websocket.Type, message []byte) {
				panic("errxxx")
			},
		})
	})
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/v4/internal/socket (interfaces: JobManager,Job,Percentable,Conn,PtyHandler,TaskManager,GorillaWs,SessionMapper)
//
// Generated by this command:
//
//	mockgen -destination ./mock_socket.go -package socket github.com/duc-cnzj/mars/v4/internal/socket JobManager,Job,Percentable,Conn,PtyHandler,TaskManager,GorillaWs,SessionMapper
//

// Package socket is a generated GoMock package.
package socket

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	websocket "github.com/duc-cnzj/mars/api/v4/websocket"
	application "github.com/duc-cnzj/mars/v4/internal/application"
	schematype "github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype"
	repo "github.com/duc-cnzj/mars/v4/internal/repo"
	gomock "go.uber.org/mock/gomock"
	remotecommand "k8s.io/client-go/tools/remotecommand"
)

// MockJobManager is a mock of JobManager interface.
type MockJobManager struct {
	ctrl     *gomock.Controller
	recorder *MockJobManagerMockRecorder
}

// MockJobManagerMockRecorder is the mock recorder for MockJobManager.
type MockJobManagerMockRecorder struct {
	mock *MockJobManager
}

// NewMockJobManager creates a new mock instance.
func NewMockJobManager(ctrl *gomock.Controller) *MockJobManager {
	mock := &MockJobManager{ctrl: ctrl}
	mock.recorder = &MockJobManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobManager) EXPECT() *MockJobManagerMockRecorder {
	return m.recorder
}

// NewJob mocks base method.
func (m *MockJobManager) NewJob(arg0 *JobInput) Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewJob", arg0)
	ret0, _ := ret[0].(Job)
	return ret0
}

// NewJob indicates an expected call of NewJob.
func (mr *MockJobManagerMockRecorder) NewJob(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewJob", reflect.TypeOf((*MockJobManager)(nil).NewJob), arg0)
}

// MockJob is a mock of Job interface.
type MockJob struct {
	ctrl     *gomock.Controller
	recorder *MockJobMockRecorder
}

// MockJobMockRecorder is the mock recorder for MockJob.
type MockJobMockRecorder struct {
	mock *MockJob
}

// NewMockJob creates a new mock instance.
func NewMockJob(ctrl *gomock.Controller) *MockJob {
	mock := &MockJob{ctrl: ctrl}
	mock.recorder = &MockJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJob) EXPECT() *MockJobMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockJob) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockJobMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockJob)(nil).Error))
}

// Finish mocks base method.
func (m *MockJob) Finish() Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(Job)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockJobMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockJob)(nil).Finish))
}

// GlobalLock mocks base method.
func (m *MockJob) GlobalLock() Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalLock")
	ret0, _ := ret[0].(Job)
	return ret0
}

// GlobalLock indicates an expected call of GlobalLock.
func (mr *MockJobMockRecorder) GlobalLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalLock", reflect.TypeOf((*MockJob)(nil).GlobalLock))
}

// ID mocks base method.
func (m *MockJob) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockJobMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockJob)(nil).ID))
}

// IsNotDryRun mocks base method.
func (m *MockJob) IsNotDryRun() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotDryRun")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotDryRun indicates an expected call of IsNotDryRun.
func (mr *MockJobMockRecorder) IsNotDryRun() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotDryRun", reflect.TypeOf((*MockJob)(nil).IsNotDryRun))
}

// LoadConfigs mocks base method.
func (m *MockJob) LoadConfigs() Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadConfigs")
	ret0, _ := ret[0].(Job)
	return ret0
}

// LoadConfigs indicates an expected call of LoadConfigs.
func (mr *MockJobMockRecorder) LoadConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadConfigs", reflect.TypeOf((*MockJob)(nil).LoadConfigs))
}

// Manifests mocks base method.
func (m *MockJob) Manifests() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifests")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Manifests indicates an expected call of Manifests.
func (mr *MockJobMockRecorder) Manifests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifests", reflect.TypeOf((*MockJob)(nil).Manifests))
}

// OnError mocks base method.
func (m *MockJob) OnError(arg0 int, arg1 func(error, func())) Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnError", arg0, arg1)
	ret0, _ := ret[0].(Job)
	return ret0
}

// OnError indicates an expected call of OnError.
func (mr *MockJobMockRecorder) OnError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockJob)(nil).OnError), arg0, arg1)
}

// OnFinally mocks base method.
func (m *MockJob) OnFinally(arg0 int, arg1 func(error, func())) Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnFinally", arg0, arg1)
	ret0, _ := ret[0].(Job)
	return ret0
}

// OnFinally indicates an expected call of OnFinally.
func (mr *MockJobMockRecorder) OnFinally(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFinally", reflect.TypeOf((*MockJob)(nil).OnFinally), arg0, arg1)
}

// OnSuccess mocks base method.
func (m *MockJob) OnSuccess(arg0 int, arg1 func(error, func())) Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnSuccess", arg0, arg1)
	ret0, _ := ret[0].(Job)
	return ret0
}

// OnSuccess indicates an expected call of OnSuccess.
func (mr *MockJobMockRecorder) OnSuccess(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSuccess", reflect.TypeOf((*MockJob)(nil).OnSuccess), arg0, arg1)
}

// Project mocks base method.
func (m *MockJob) Project() *repo.Project {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project")
	ret0, _ := ret[0].(*repo.Project)
	return ret0
}

// Project indicates an expected call of Project.
func (mr *MockJobMockRecorder) Project() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockJob)(nil).Project))
}

// Run mocks base method.
func (m *MockJob) Run(arg0 context.Context) Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(Job)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockJobMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockJob)(nil).Run), arg0)
}

// Stop mocks base method.
func (m *MockJob) Stop(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop.
func (mr *MockJobMockRecorder) Stop(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJob)(nil).Stop), arg0)
}

// Validate mocks base method.
func (m *MockJob) Validate() Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(Job)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockJobMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockJob)(nil).Validate))
}

// MockPercentable is a mock of Percentable interface.
type MockPercentable struct {
	ctrl     *gomock.Controller
	recorder *MockPercentableMockRecorder
}

// MockPercentableMockRecorder is the mock recorder for MockPercentable.
type MockPercentableMockRecorder struct {
	mock *MockPercentable
}

// NewMockPercentable creates a new mock instance.
func NewMockPercentable(ctrl *gomock.Controller) *MockPercentable {
	mock := &MockPercentable{ctrl: ctrl}
	mock.recorder = &MockPercentableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPercentable) EXPECT() *MockPercentableMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPercentable) Add() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add")
}

// Add indicates an expected call of Add.
func (mr *MockPercentableMockRecorder) Add() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPercentable)(nil).Add))
}

// Current mocks base method.
func (m *MockPercentable) Current() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockPercentableMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockPercentable)(nil).Current))
}

// To mocks base method.
func (m *MockPercentable) To(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "To", arg0)
}

// To indicates an expected call of To.
func (mr *MockPercentableMockRecorder) To(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "To", reflect.TypeOf((*MockPercentable)(nil).To), arg0)
}

// MockConn is a mock of Conn interface.
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn.
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance.
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockConn) AddTask(arg0 string, arg1 func(error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask.
func (mr *MockConnMockRecorder) AddTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockConn)(nil).AddTask), arg0, arg1)
}

// Close mocks base method.
func (m *MockConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}

// CloseAndClean mocks base method.
func (m *MockConn) CloseAndClean(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndClean", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseAndClean indicates an expected call of CloseAndClean.
func (mr *MockConnMockRecorder) CloseAndClean(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndClean", reflect.TypeOf((*MockConn)(nil).CloseAndClean), arg0)
}

// ClosePty mocks base method.
func (m *MockConn) ClosePty(arg0 context.Context, arg1 string, arg2 uint32, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClosePty", arg0, arg1, arg2, arg3)
}

// ClosePty indicates an expected call of ClosePty.
func (mr *MockConnMockRecorder) ClosePty(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePty", reflect.TypeOf((*MockConn)(nil).ClosePty), arg0, arg1, arg2, arg3)
}

// GetPtyHandler mocks base method.
func (m *MockConn) GetPtyHandler(arg0 string) (PtyHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPtyHandler", arg0)
	ret0, _ := ret[0].(PtyHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPtyHandler indicates an expected call of GetPtyHandler.
func (mr *MockConnMockRecorder) GetPtyHandler(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPtyHandler", reflect.TypeOf((*MockConn)(nil).GetPtyHandler), arg0)
}

// GetUser mocks base method.
func (m *MockConn) GetUser() *schematype.UserInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].(*schematype.UserInfo)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockConnMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockConn)(nil).GetUser))
}

// ID mocks base method.
func (m *MockConn) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockConnMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockConn)(nil).ID))
}

// NextWriter mocks base method.
func (m *MockConn) NextWriter(arg0 int) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextWriter", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextWriter indicates an expected call of NextWriter.
func (mr *MockConnMockRecorder) NextWriter(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWriter", reflect.TypeOf((*MockConn)(nil).NextWriter), arg0)
}

// PubSub mocks base method.
func (m *MockConn) PubSub() application.PubSub {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSub")
	ret0, _ := ret[0].(application.PubSub)
	return ret0
}

// PubSub indicates an expected call of PubSub.
func (mr *MockConnMockRecorder) PubSub() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSub", reflect.TypeOf((*MockConn)(nil).PubSub))
}

// ReadMessage mocks base method.
func (m *MockConn) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockConnMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockConn)(nil).ReadMessage))
}

// RemoveTask mocks base method.
func (m *MockConn) RemoveTask(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTask", arg0)
}

// RemoveTask indicates an expected call of RemoveTask.
func (mr *MockConnMockRecorder) RemoveTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockConn)(nil).RemoveTask), arg0)
}

// RunTask mocks base method.
func (m *MockConn) RunTask(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTask indicates an expected call of RunTask.
func (mr *MockConnMockRecorder) RunTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTask", reflect.TypeOf((*MockConn)(nil).RunTask), arg0)
}

// SetPongHandler mocks base method.
func (m *MockConn) SetPongHandler(arg0 func(string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPongHandler", arg0)
}

// SetPongHandler indicates an expected call of SetPongHandler.
func (mr *MockConnMockRecorder) SetPongHandler(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPongHandler", reflect.TypeOf((*MockConn)(nil).SetPongHandler), arg0)
}

// SetPtyHandler mocks base method.
func (m *MockConn) SetPtyHandler(arg0 string, arg1 PtyHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPtyHandler", arg0, arg1)
}

// SetPtyHandler indicates an expected call of SetPtyHandler.
func (mr *MockConnMockRecorder) SetPtyHandler(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPtyHandler", reflect.TypeOf((*MockConn)(nil).SetPtyHandler), arg0, arg1)
}

// SetReadDeadline mocks base method.
func (m *MockConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockConnMockRecorder) SetReadDeadline(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConn)(nil).SetReadDeadline), arg0)
}

// SetReadLimit mocks base method.
func (m *MockConn) SetReadLimit(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadLimit", arg0)
}

// SetReadLimit indicates an expected call of SetReadLimit.
func (mr *MockConnMockRecorder) SetReadLimit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLimit", reflect.TypeOf((*MockConn)(nil).SetReadLimit), arg0)
}

// SetUser mocks base method.
func (m *MockConn) SetUser(arg0 *schematype.UserInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUser", arg0)
}

// SetUser indicates an expected call of SetUser.
func (mr *MockConnMockRecorder) SetUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUser", reflect.TypeOf((*MockConn)(nil).SetUser), arg0)
}

// SetWriteDeadline mocks base method.
func (m *MockConn) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockConnMockRecorder) SetWriteDeadline(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockConn)(nil).SetWriteDeadline), arg0)
}

// UID mocks base method.
func (m *MockConn) UID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UID indicates an expected call of UID.
func (mr *MockConnMockRecorder) UID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockConn)(nil).UID))
}

// WriteMessage mocks base method.
func (m *MockConn) WriteMessage(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockConnMockRecorder) WriteMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockConn)(nil).WriteMessage), arg0, arg1)
}

// MockPtyHandler is a mock of PtyHandler interface.
type MockPtyHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPtyHandlerMockRecorder
}

// MockPtyHandlerMockRecorder is the mock recorder for MockPtyHandler.
type MockPtyHandlerMockRecorder struct {
	mock *MockPtyHandler
}

// NewMockPtyHandler creates a new mock instance.
func NewMockPtyHandler(ctrl *gomock.Controller) *MockPtyHandler {
	mock := &MockPtyHandler{ctrl: ctrl}
	mock.recorder = &MockPtyHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPtyHandler) EXPECT() *MockPtyHandlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPtyHandler) Close(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPtyHandlerMockRecorder) Close(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPtyHandler)(nil).Close), arg0, arg1)
}

// Cols mocks base method.
func (m *MockPtyHandler) Cols() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cols")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Cols indicates an expected call of Cols.
func (mr *MockPtyHandlerMockRecorder) Cols() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cols", reflect.TypeOf((*MockPtyHandler)(nil).Cols))
}

// Container mocks base method.
func (m *MockPtyHandler) Container() *repo.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Container")
	ret0, _ := ret[0].(*repo.Container)
	return ret0
}

// Container indicates an expected call of Container.
func (mr *MockPtyHandlerMockRecorder) Container() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Container", reflect.TypeOf((*MockPtyHandler)(nil).Container))
}

// IsClosed mocks base method.
func (m *MockPtyHandler) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockPtyHandlerMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockPtyHandler)(nil).IsClosed))
}

// Next mocks base method.
func (m *MockPtyHandler) Next() *remotecommand.TerminalSize {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*remotecommand.TerminalSize)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockPtyHandlerMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPtyHandler)(nil).Next))
}

// Read mocks base method.
func (m *MockPtyHandler) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPtyHandlerMockRecorder) Read(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPtyHandler)(nil).Read), arg0)
}

// Recorder mocks base method.
func (m *MockPtyHandler) Recorder() repo.Recorder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recorder")
	ret0, _ := ret[0].(repo.Recorder)
	return ret0
}

// Recorder indicates an expected call of Recorder.
func (mr *MockPtyHandlerMockRecorder) Recorder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recorder", reflect.TypeOf((*MockPtyHandler)(nil).Recorder))
}

// ResetTerminalRowCol mocks base method.
func (m *MockPtyHandler) ResetTerminalRowCol(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetTerminalRowCol", arg0)
}

// ResetTerminalRowCol indicates an expected call of ResetTerminalRowCol.
func (mr *MockPtyHandlerMockRecorder) ResetTerminalRowCol(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTerminalRowCol", reflect.TypeOf((*MockPtyHandler)(nil).ResetTerminalRowCol), arg0)
}

// Resize mocks base method.
func (m *MockPtyHandler) Resize(arg0 remotecommand.TerminalSize) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resize indicates an expected call of Resize.
func (mr *MockPtyHandlerMockRecorder) Resize(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockPtyHandler)(nil).Resize), arg0)
}

// Rows mocks base method.
func (m *MockPtyHandler) Rows() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rows")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Rows indicates an expected call of Rows.
func (mr *MockPtyHandlerMockRecorder) Rows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rows", reflect.TypeOf((*MockPtyHandler)(nil).Rows))
}

// Send mocks base method.
func (m *MockPtyHandler) Send(arg0 context.Context, arg1 *websocket.TerminalMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockPtyHandlerMockRecorder) Send(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPtyHandler)(nil).Send), arg0, arg1)
}

// SetShell mocks base method.
func (m *MockPtyHandler) SetShell(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShell", arg0)
}

// SetShell indicates an expected call of SetShell.
func (mr *MockPtyHandlerMockRecorder) SetShell(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShell", reflect.TypeOf((*MockPtyHandler)(nil).SetShell), arg0)
}

// Toast mocks base method.
func (m *MockPtyHandler) Toast(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Toast", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Toast indicates an expected call of Toast.
func (mr *MockPtyHandlerMockRecorder) Toast(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Toast", reflect.TypeOf((*MockPtyHandler)(nil).Toast), arg0)
}

// Write mocks base method.
func (m *MockPtyHandler) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockPtyHandlerMockRecorder) Write(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockPtyHandler)(nil).Write), arg0)
}

// MockTaskManager is a mock of TaskManager interface.
type MockTaskManager struct {
	ctrl     *gomock.Controller
	recorder *MockTaskManagerMockRecorder
}

// MockTaskManagerMockRecorder is the mock recorder for MockTaskManager.
type MockTaskManagerMockRecorder struct {
	mock *MockTaskManager
}

// NewMockTaskManager creates a new mock instance.
func NewMockTaskManager(ctrl *gomock.Controller) *MockTaskManager {
	mock := &MockTaskManager{ctrl: ctrl}
	mock.recorder = &MockTaskManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskManager) EXPECT() *MockTaskManagerMockRecorder {
	return m.recorder
}

// Has mocks base method.
func (m *MockTaskManager) Has(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockTaskManagerMockRecorder) Has(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTaskManager)(nil).Has), arg0)
}

// Register mocks base method.
func (m *MockTaskManager) Register(arg0 string, arg1 func(error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockTaskManagerMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockTaskManager)(nil).Register), arg0, arg1)
}

// Remove mocks base method.
func (m *MockTaskManager) Remove(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockTaskManagerMockRecorder) Remove(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTaskManager)(nil).Remove), arg0)
}

// Stop mocks base method.
func (m *MockTaskManager) Stop(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop.
func (mr *MockTaskManagerMockRecorder) Stop(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTaskManager)(nil).Stop), arg0)
}

// StopAll mocks base method.
func (m *MockTaskManager) StopAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopAll")
}

// StopAll indicates an expected call of StopAll.
func (mr *MockTaskManagerMockRecorder) StopAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAll", reflect.TypeOf((*MockTaskManager)(nil).StopAll))
}

// MockGorillaWs is a mock of GorillaWs interface.
type MockGorillaWs struct {
	ctrl     *gomock.Controller
	recorder *MockGorillaWsMockRecorder
}

// MockGorillaWsMockRecorder is the mock recorder for MockGorillaWs.
type MockGorillaWsMockRecorder struct {
	mock *MockGorillaWs
}

// NewMockGorillaWs creates a new mock instance.
func NewMockGorillaWs(ctrl *gomock.Controller) *MockGorillaWs {
	mock := &MockGorillaWs{ctrl: ctrl}
	mock.recorder = &MockGorillaWsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGorillaWs) EXPECT() *MockGorillaWsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGorillaWs) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGorillaWsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGorillaWs)(nil).Close))
}

// NextWriter mocks base method.
func (m *MockGorillaWs) NextWriter(arg0 int) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextWriter", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextWriter indicates an expected call of NextWriter.
func (mr *MockGorillaWsMockRecorder) NextWriter(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWriter", reflect.TypeOf((*MockGorillaWs)(nil).NextWriter), arg0)
}

// ReadMessage mocks base method.
func (m *MockGorillaWs) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockGorillaWsMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockGorillaWs)(nil).ReadMessage))
}

// SetPongHandler mocks base method.
func (m *MockGorillaWs) SetPongHandler(arg0 func(string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPongHandler", arg0)
}

// SetPongHandler indicates an expected call of SetPongHandler.
func (mr *MockGorillaWsMockRecorder) SetPongHandler(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPongHandler", reflect.TypeOf((*MockGorillaWs)(nil).SetPongHandler), arg0)
}

// SetReadDeadline mocks base method.
func (m *MockGorillaWs) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockGorillaWsMockRecorder) SetReadDeadline(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockGorillaWs)(nil).SetReadDeadline), arg0)
}

// SetReadLimit mocks base method.
func (m *MockGorillaWs) SetReadLimit(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadLimit", arg0)
}

// SetReadLimit indicates an expected call of SetReadLimit.
func (mr *MockGorillaWsMockRecorder) SetReadLimit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLimit", reflect.TypeOf((*MockGorillaWs)(nil).SetReadLimit), arg0)
}

// SetWriteDeadline mocks base method.
func (m *MockGorillaWs) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockGorillaWsMockRecorder) SetWriteDeadline(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockGorillaWs)(nil).SetWriteDeadline), arg0)
}

// WriteMessage mocks base method.
func (m *MockGorillaWs) WriteMessage(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockGorillaWsMockRecorder) WriteMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockGorillaWs)(nil).WriteMessage), arg0, arg1)
}

// MockSessionMapper is a mock of SessionMapper interface.
type MockSessionMapper struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMapperMockRecorder
}

// MockSessionMapperMockRecorder is the mock recorder for MockSessionMapper.
type MockSessionMapperMockRecorder struct {
	mock *MockSessionMapper
}

// NewMockSessionMapper creates a new mock instance.
func NewMockSessionMapper(ctrl *gomock.Controller) *MockSessionMapper {
	mock := &MockSessionMapper{ctrl: ctrl}
	mock.recorder = &MockSessionMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionMapper) EXPECT() *MockSessionMapperMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSessionMapper) Close(arg0 context.Context, arg1 string, arg2 uint32, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", arg0, arg1, arg2, arg3)
}

// Close indicates an expected call of Close.
func (mr *MockSessionMapperMockRecorder) Close(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSessionMapper)(nil).Close), arg0, arg1, arg2, arg3)
}

// CloseAll mocks base method.
func (m *MockSessionMapper) CloseAll(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseAll", arg0)
}

// CloseAll indicates an expected call of CloseAll.
func (mr *MockSessionMapperMockRecorder) CloseAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAll", reflect.TypeOf((*MockSessionMapper)(nil).CloseAll), arg0)
}

// Get mocks base method.
func (m *MockSessionMapper) Get(arg0 string) (PtyHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(PtyHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionMapperMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionMapper)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockSessionMapper) Set(arg0 string, arg1 PtyHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockSessionMapperMockRecorder) Set(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSessionMapper)(nil).Set), arg0, arg1)
}

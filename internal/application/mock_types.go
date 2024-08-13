// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/v4/internal/application (interfaces: PluginManger,Picture,Project)
//
// Generated by this command:
//
//	mockgen -destination ./mock_types.go -package application github.com/duc-cnzj/mars/v4/internal/application PluginManger,Picture,Project
//

// Package application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPluginManger is a mock of PluginManger interface.
type MockPluginManger struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMangerMockRecorder
}

// MockPluginMangerMockRecorder is the mock recorder for MockPluginManger.
type MockPluginMangerMockRecorder struct {
	mock *MockPluginManger
}

// NewMockPluginManger creates a new mock instance.
func NewMockPluginManger(ctrl *gomock.Controller) *MockPluginManger {
	mock := &MockPluginManger{ctrl: ctrl}
	mock.recorder = &MockPluginMangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPluginManger) EXPECT() *MockPluginMangerMockRecorder {
	return m.recorder
}

// Domain mocks base method.
func (m *MockPluginManger) Domain() DomainManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Domain")
	ret0, _ := ret[0].(DomainManager)
	return ret0
}

// Domain indicates an expected call of Domain.
func (mr *MockPluginMangerMockRecorder) Domain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Domain", reflect.TypeOf((*MockPluginManger)(nil).Domain))
}

// GetPlugins mocks base method.
func (m *MockPluginManger) GetPlugins() map[string]Plugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlugins")
	ret0, _ := ret[0].(map[string]Plugin)
	return ret0
}

// GetPlugins indicates an expected call of GetPlugins.
func (mr *MockPluginMangerMockRecorder) GetPlugins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlugins", reflect.TypeOf((*MockPluginManger)(nil).GetPlugins))
}

// Git mocks base method.
func (m *MockPluginManger) Git() GitServer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Git")
	ret0, _ := ret[0].(GitServer)
	return ret0
}

// Git indicates an expected call of Git.
func (mr *MockPluginMangerMockRecorder) Git() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Git", reflect.TypeOf((*MockPluginManger)(nil).Git))
}

// Load mocks base method.
func (m *MockPluginManger) Load(arg0 App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockPluginMangerMockRecorder) Load(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockPluginManger)(nil).Load), arg0)
}

// Picture mocks base method.
func (m *MockPluginManger) Picture() Picture {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Picture")
	ret0, _ := ret[0].(Picture)
	return ret0
}

// Picture indicates an expected call of Picture.
func (mr *MockPluginMangerMockRecorder) Picture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Picture", reflect.TypeOf((*MockPluginManger)(nil).Picture))
}

// Ws mocks base method.
func (m *MockPluginManger) Ws() WsSender {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ws")
	ret0, _ := ret[0].(WsSender)
	return ret0
}

// Ws indicates an expected call of Ws.
func (mr *MockPluginMangerMockRecorder) Ws() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ws", reflect.TypeOf((*MockPluginManger)(nil).Ws))
}

// MockPicture is a mock of Picture interface.
type MockPicture struct {
	ctrl     *gomock.Controller
	recorder *MockPictureMockRecorder
}

// MockPictureMockRecorder is the mock recorder for MockPicture.
type MockPictureMockRecorder struct {
	mock *MockPicture
}

// NewMockPicture creates a new mock instance.
func NewMockPicture(ctrl *gomock.Controller) *MockPicture {
	mock := &MockPicture{ctrl: ctrl}
	mock.recorder = &MockPictureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPicture) EXPECT() *MockPictureMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockPicture) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockPictureMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockPicture)(nil).Destroy))
}

// Get mocks base method.
func (m *MockPicture) Get(arg0 context.Context, arg1 bool) (*PictureItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*PictureItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPictureMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPicture)(nil).Get), arg0, arg1)
}

// Initialize mocks base method.
func (m *MockPicture) Initialize(arg0 App, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockPictureMockRecorder) Initialize(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockPicture)(nil).Initialize), arg0, arg1)
}

// Name mocks base method.
func (m *MockPicture) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPictureMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPicture)(nil).Name))
}

// MockProject is a mock of Project interface.
type MockProject struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMockRecorder
}

// MockProjectMockRecorder is the mock recorder for MockProject.
type MockProjectMockRecorder struct {
	mock *MockProject
}

// NewMockProject creates a new mock instance.
func NewMockProject(ctrl *gomock.Controller) *MockProject {
	mock := &MockProject{ctrl: ctrl}
	mock.recorder = &MockProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProject) EXPECT() *MockProjectMockRecorder {
	return m.recorder
}

// GetAvatarURL mocks base method.
func (m *MockProject) GetAvatarURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatarURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAvatarURL indicates an expected call of GetAvatarURL.
func (mr *MockProjectMockRecorder) GetAvatarURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatarURL", reflect.TypeOf((*MockProject)(nil).GetAvatarURL))
}

// GetDefaultBranch mocks base method.
func (m *MockProject) GetDefaultBranch() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultBranch")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDefaultBranch indicates an expected call of GetDefaultBranch.
func (mr *MockProjectMockRecorder) GetDefaultBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultBranch", reflect.TypeOf((*MockProject)(nil).GetDefaultBranch))
}

// GetDescription mocks base method.
func (m *MockProject) GetDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription.
func (mr *MockProjectMockRecorder) GetDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescription", reflect.TypeOf((*MockProject)(nil).GetDescription))
}

// GetID mocks base method.
func (m *MockProject) GetID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockProjectMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockProject)(nil).GetID))
}

// GetName mocks base method.
func (m *MockProject) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockProjectMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockProject)(nil).GetName))
}

// GetPath mocks base method.
func (m *MockProject) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockProjectMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockProject)(nil).GetPath))
}

// GetWebURL mocks base method.
func (m *MockProject) GetWebURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWebURL indicates an expected call of GetWebURL.
func (mr *MockProjectMockRecorder) GetWebURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebURL", reflect.TypeOf((*MockProject)(nil).GetWebURL))
}

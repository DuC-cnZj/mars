// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/internal/contracts (interfaces: DispatcherInterface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	contracts "github.com/duc-cnzj/mars/internal/contracts"
	gomock "github.com/golang/mock/gomock"
)

// MockDispatcherInterface is a mock of DispatcherInterface interface.
type MockDispatcherInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherInterfaceMockRecorder
}

// MockDispatcherInterfaceMockRecorder is the mock recorder for MockDispatcherInterface.
type MockDispatcherInterfaceMockRecorder struct {
	mock *MockDispatcherInterface
}

// NewMockDispatcherInterface creates a new mock instance.
func NewMockDispatcherInterface(ctrl *gomock.Controller) *MockDispatcherInterface {
	mock := &MockDispatcherInterface{ctrl: ctrl}
	mock.recorder = &MockDispatcherInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDispatcherInterface) EXPECT() *MockDispatcherInterfaceMockRecorder {
	return m.recorder
}

// Dispatch mocks base method.
func (m *MockDispatcherInterface) Dispatch(arg0 contracts.Event, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockDispatcherInterfaceMockRecorder) Dispatch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockDispatcherInterface)(nil).Dispatch), arg0, arg1)
}

// Forget mocks base method.
func (m *MockDispatcherInterface) Forget(arg0 contracts.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Forget", arg0)
}

// Forget indicates an expected call of Forget.
func (mr *MockDispatcherInterfaceMockRecorder) Forget(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forget", reflect.TypeOf((*MockDispatcherInterface)(nil).Forget), arg0)
}

// GetListeners mocks base method.
func (m *MockDispatcherInterface) GetListeners(arg0 contracts.Event) []contracts.Listener {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListeners", arg0)
	ret0, _ := ret[0].([]contracts.Listener)
	return ret0
}

// GetListeners indicates an expected call of GetListeners.
func (mr *MockDispatcherInterfaceMockRecorder) GetListeners(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListeners", reflect.TypeOf((*MockDispatcherInterface)(nil).GetListeners), arg0)
}

// HasListeners mocks base method.
func (m *MockDispatcherInterface) HasListeners(arg0 contracts.Event) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasListeners", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasListeners indicates an expected call of HasListeners.
func (mr *MockDispatcherInterfaceMockRecorder) HasListeners(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasListeners", reflect.TypeOf((*MockDispatcherInterface)(nil).HasListeners), arg0)
}

// Listen mocks base method.
func (m *MockDispatcherInterface) Listen(arg0 contracts.Event, arg1 contracts.Listener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Listen", arg0, arg1)
}

// Listen indicates an expected call of Listen.
func (mr *MockDispatcherInterfaceMockRecorder) Listen(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockDispatcherInterface)(nil).Listen), arg0, arg1)
}

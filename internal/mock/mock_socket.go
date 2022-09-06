// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/internal/contracts (interfaces: CancelSignaler)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCancelSignaler is a mock of CancelSignaler interface.
type MockCancelSignaler struct {
	ctrl     *gomock.Controller
	recorder *MockCancelSignalerMockRecorder
}

// MockCancelSignalerMockRecorder is the mock recorder for MockCancelSignaler.
type MockCancelSignalerMockRecorder struct {
	mock *MockCancelSignaler
}

// NewMockCancelSignaler creates a new mock instance.
func NewMockCancelSignaler(ctrl *gomock.Controller) *MockCancelSignaler {
	mock := &MockCancelSignaler{ctrl: ctrl}
	mock.recorder = &MockCancelSignalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCancelSignaler) EXPECT() *MockCancelSignalerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCancelSignaler) Add(arg0 string, arg1 func(error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCancelSignalerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCancelSignaler)(nil).Add), arg0, arg1)
}

// Cancel mocks base method.
func (m *MockCancelSignaler) Cancel(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel", arg0)
}

// Cancel indicates an expected call of Cancel.
func (mr *MockCancelSignalerMockRecorder) Cancel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockCancelSignaler)(nil).Cancel), arg0)
}

// CancelAll mocks base method.
func (m *MockCancelSignaler) CancelAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelAll")
}

// CancelAll indicates an expected call of CancelAll.
func (mr *MockCancelSignalerMockRecorder) CancelAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAll", reflect.TypeOf((*MockCancelSignaler)(nil).CancelAll))
}

// Has mocks base method.
func (m *MockCancelSignaler) Has(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockCancelSignalerMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCancelSignaler)(nil).Has), arg0)
}

// Remove mocks base method.
func (m *MockCancelSignaler) Remove(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockCancelSignalerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCancelSignaler)(nil).Remove), arg0)
}

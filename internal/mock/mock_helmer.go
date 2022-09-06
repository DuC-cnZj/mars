// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/internal/contracts (interfaces: Helmer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	types "github.com/duc-cnzj/mars-client/v4/types"
	contracts "github.com/duc-cnzj/mars/internal/contracts"
	gomock "github.com/golang/mock/gomock"
	chart "helm.sh/helm/v3/pkg/chart"
	values "helm.sh/helm/v3/pkg/cli/values"
	release "helm.sh/helm/v3/pkg/release"
)

// MockHelmer is a mock of Helmer interface.
type MockHelmer struct {
	ctrl     *gomock.Controller
	recorder *MockHelmerMockRecorder
}

// MockHelmerMockRecorder is the mock recorder for MockHelmer.
type MockHelmerMockRecorder struct {
	mock *MockHelmer
}

// NewMockHelmer creates a new mock instance.
func NewMockHelmer(ctrl *gomock.Controller) *MockHelmer {
	mock := &MockHelmer{ctrl: ctrl}
	mock.recorder = &MockHelmerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelmer) EXPECT() *MockHelmerMockRecorder {
	return m.recorder
}

// PackageChart mocks base method.
func (m *MockHelmer) PackageChart(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackageChart", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackageChart indicates an expected call of PackageChart.
func (mr *MockHelmerMockRecorder) PackageChart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackageChart", reflect.TypeOf((*MockHelmer)(nil).PackageChart), arg0, arg1)
}

// ReleaseStatus mocks base method.
func (m *MockHelmer) ReleaseStatus(arg0, arg1 string) types.Deploy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseStatus", arg0, arg1)
	ret0, _ := ret[0].(types.Deploy)
	return ret0
}

// ReleaseStatus indicates an expected call of ReleaseStatus.
func (mr *MockHelmerMockRecorder) ReleaseStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseStatus", reflect.TypeOf((*MockHelmer)(nil).ReleaseStatus), arg0, arg1)
}

// Rollback mocks base method.
func (m *MockHelmer) Rollback(arg0, arg1 string, arg2 bool, arg3 contracts.LogFn, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockHelmerMockRecorder) Rollback(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockHelmer)(nil).Rollback), arg0, arg1, arg2, arg3, arg4)
}

// Uninstall mocks base method.
func (m *MockHelmer) Uninstall(arg0, arg1 string, arg2 contracts.LogFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockHelmerMockRecorder) Uninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockHelmer)(nil).Uninstall), arg0, arg1, arg2)
}

// UpgradeOrInstall mocks base method.
func (m *MockHelmer) UpgradeOrInstall(arg0 context.Context, arg1, arg2 string, arg3 *chart.Chart, arg4 *values.Options, arg5 contracts.WrapLogFn, arg6 bool, arg7 int64, arg8 bool) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeOrInstall", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeOrInstall indicates an expected call of UpgradeOrInstall.
func (mr *MockHelmerMockRecorder) UpgradeOrInstall(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeOrInstall", reflect.TypeOf((*MockHelmer)(nil).UpgradeOrInstall), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

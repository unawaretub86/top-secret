package mocks_test

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockMessagePort struct {
	ctrl     *gomock.Controller
	recorder *MockMessagePortMockRecorder
}

type MockMessagePortMockRecorder struct {
	mock *MockMessagePort
}

func NewMockMessagePort(ctrl *gomock.Controller) *MockMessagePort {
	mock := &MockMessagePort{ctrl: ctrl}
	mock.recorder = &MockMessagePortMockRecorder{mock}
	return mock
}

func (m *MockMessagePort) EXPECT() *MockMessagePortMockRecorder {
	return m.recorder
}

func (m *MockMessagePort) GetMessage(requestID string, messages ...[]string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", messages)
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockMessagePortMockRecorder) GetMessage(requestID string, messages ...[]string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockMessagePort)(nil).GetMessage), messages)
}

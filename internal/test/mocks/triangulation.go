package mocks_test

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockTriangulationPort struct {
	ctrl     *gomock.Controller
	recorder *MockTriangulationPortMockRecorder
}

type MockTriangulationPortMockRecorder struct {
	mock *MockTriangulationPort
}

func NewMockTriangulationPort(ctrl *gomock.Controller) *MockTriangulationPort {
	mock := &MockTriangulationPort{ctrl: ctrl}
	mock.recorder = &MockTriangulationPortMockRecorder{mock}
	return mock
}

func (m *MockTriangulationPort) EXPECT() *MockTriangulationPortMockRecorder {
	return m.recorder
}

func (m *MockTriangulationPort) GetLocation(requestID string, distances ...float32) (float32, float32) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", distances)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(float32)
	return ret0, ret1
}

func (mr *MockTriangulationPortMockRecorder) GetLocation(requestID string, distances ...float32) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockTriangulationPort)(nil).GetLocation), distances)
}

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

// crea una nueva instancia del mock.
// recibe un controlador de gomock como argumento y devuelve un mock.
func NewMockTriangulationPort(ctrl *gomock.Controller) *MockTriangulationPort {
	mock := &MockTriangulationPort{ctrl: ctrl}
	mock.recorder = &MockTriangulationPortMockRecorder{mock}
	return mock
}

// EXPECT() es para registrar llamadas a métodos en el mock.
func (m *MockTriangulationPort) EXPECT() *MockTriangulationPortMockRecorder {
	return m.recorder
}

// Recibe un argumento requestID  y distances
// y devuelve una respuesta simulada.
func (m *MockTriangulationPort) GetLocation(requestID string, distances ...float32) (float32, float32) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", distances)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(float32)
	return ret0, ret1
}

// GetLocation  se utiliza para obtener llamadas al método GetLocation en el mock.
func (mr *MockTriangulationPortMockRecorder) GetLocation(requestID string, distances ...float32) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockTriangulationPort)(nil).GetLocation), distances)
}

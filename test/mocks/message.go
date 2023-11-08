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

// crea una nueva instancia del mock.
// recibe un controlador de gomock como argumento y devuelve un mock.
func NewMockMessagePort(ctrl *gomock.Controller) *MockMessagePort {
	mock := &MockMessagePort{ctrl: ctrl}
	mock.recorder = &MockMessagePortMockRecorder{mock}
	return mock
}

// EXPECT() es para registrar llamadas a métodos en el mock.
func (m *MockMessagePort) EXPECT() *MockMessagePortMockRecorder {
	return m.recorder
}

// Recibe un argumento requestID  y distances
// y devuelve una respuesta simulada.
func (m *MockMessagePort) GetMessage(requestID string, messages ...[]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", messages)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage  se utiliza para obtener llamadas al método GetMessage en el mock.
func (mr *MockMessagePortMockRecorder) GetMessage(requestID string, messages ...[]string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockMessagePort)(nil).GetMessage), messages)
}

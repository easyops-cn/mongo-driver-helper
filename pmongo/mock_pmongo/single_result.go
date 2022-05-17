// Code generated by MockGen. DO NOT EDIT.
// Source: pmongo/single_result.go

// Package mock_pmongo is a generated GoMock package.
package mock_pmongo

import (
	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
	reflect "reflect"
)

// MockSingleResultInterface is a mock of SingleResultInterface interface
type MockSingleResultInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSingleResultInterfaceMockRecorder
}

// MockSingleResultInterfaceMockRecorder is the mock recorder for MockSingleResultInterface
type MockSingleResultInterfaceMockRecorder struct {
	mock *MockSingleResultInterface
}

// NewMockSingleResultInterface creates a new mock instance
func NewMockSingleResultInterface(ctrl *gomock.Controller) *MockSingleResultInterface {
	mock := &MockSingleResultInterface{ctrl: ctrl}
	mock.recorder = &MockSingleResultInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSingleResultInterface) EXPECT() *MockSingleResultInterfaceMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockSingleResultInterface) Decode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockSingleResultInterfaceMockRecorder) Decode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockSingleResultInterface)(nil).Decode), v)
}

// DecodeBytes mocks base method
func (m *MockSingleResultInterface) DecodeBytes() (bson.Raw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeBytes")
	ret0, _ := ret[0].(bson.Raw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeBytes indicates an expected call of DecodeBytes
func (mr *MockSingleResultInterfaceMockRecorder) DecodeBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeBytes", reflect.TypeOf((*MockSingleResultInterface)(nil).DecodeBytes))
}

// Err mocks base method
func (m *MockSingleResultInterface) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockSingleResultInterfaceMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSingleResultInterface)(nil).Err))
}
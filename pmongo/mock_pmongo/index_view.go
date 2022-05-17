// Code generated by MockGen. DO NOT EDIT.
// Source: pmongo/index_view.go

// Package mock_pmongo is a generated GoMock package.
package mock_pmongo

import (
	context "context"
	pmongo "github.com/easyops-cn/mongo-driver-helper/pmongo"
	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	reflect "reflect"
)

// MockIndexViewInterface is a mock of IndexViewInterface interface
type MockIndexViewInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIndexViewInterfaceMockRecorder
}

// MockIndexViewInterfaceMockRecorder is the mock recorder for MockIndexViewInterface
type MockIndexViewInterfaceMockRecorder struct {
	mock *MockIndexViewInterface
}

// NewMockIndexViewInterface creates a new mock instance
func NewMockIndexViewInterface(ctrl *gomock.Controller) *MockIndexViewInterface {
	mock := &MockIndexViewInterface{ctrl: ctrl}
	mock.recorder = &MockIndexViewInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexViewInterface) EXPECT() *MockIndexViewInterfaceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockIndexViewInterface) List(ctx context.Context, opts ...*options.ListIndexesOptions) (pmongo.CursorInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(pmongo.CursorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockIndexViewInterfaceMockRecorder) List(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIndexViewInterface)(nil).List), varargs...)
}

// CreateOne mocks base method
func (m *MockIndexViewInterface) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, model}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOne", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOne indicates an expected call of CreateOne
func (mr *MockIndexViewInterfaceMockRecorder) CreateOne(ctx, model interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, model}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockIndexViewInterface)(nil).CreateOne), varargs...)
}

// CreateMany mocks base method
func (m *MockIndexViewInterface) CreateMany(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, models}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMany", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMany indicates an expected call of CreateMany
func (mr *MockIndexViewInterfaceMockRecorder) CreateMany(ctx, models interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, models}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockIndexViewInterface)(nil).CreateMany), varargs...)
}

// DropOne mocks base method
func (m *MockIndexViewInterface) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropOne", varargs...)
	ret0, _ := ret[0].(bson.Raw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DropOne indicates an expected call of DropOne
func (mr *MockIndexViewInterfaceMockRecorder) DropOne(ctx, name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropOne", reflect.TypeOf((*MockIndexViewInterface)(nil).DropOne), varargs...)
}

// DropAll mocks base method
func (m *MockIndexViewInterface) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropAll", varargs...)
	ret0, _ := ret[0].(bson.Raw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DropAll indicates an expected call of DropAll
func (mr *MockIndexViewInterfaceMockRecorder) DropAll(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropAll", reflect.TypeOf((*MockIndexViewInterface)(nil).DropAll), varargs...)
}
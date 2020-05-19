// Code generated by MockGen. DO NOT EDIT.
// Source: pmongo/gridfs.go

// Package mock_pmongo is a generated GoMock package.
package mock_pmongo

import (
	pmongo "github.com/easyops-cn/mongo-driver-helper/pmongo"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	options "go.mongodb.org/mongo-driver/mongo/options"
	io "io"
	reflect "reflect"
	time "time"
)

// MockDownloadStreamInterface is a mock of DownloadStreamInterface interface
type MockDownloadStreamInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadStreamInterfaceMockRecorder
}

// MockDownloadStreamInterfaceMockRecorder is the mock recorder for MockDownloadStreamInterface
type MockDownloadStreamInterfaceMockRecorder struct {
	mock *MockDownloadStreamInterface
}

// NewMockDownloadStreamInterface creates a new mock instance
func NewMockDownloadStreamInterface(ctrl *gomock.Controller) *MockDownloadStreamInterface {
	mock := &MockDownloadStreamInterface{ctrl: ctrl}
	mock.recorder = &MockDownloadStreamInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDownloadStreamInterface) EXPECT() *MockDownloadStreamInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDownloadStreamInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDownloadStreamInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDownloadStreamInterface)(nil).Close))
}

// Read mocks base method
func (m *MockDownloadStreamInterface) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockDownloadStreamInterfaceMockRecorder) Read(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDownloadStreamInterface)(nil).Read), p)
}

// SetReadDeadline mocks base method
func (m *MockDownloadStreamInterface) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockDownloadStreamInterfaceMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockDownloadStreamInterface)(nil).SetReadDeadline), t)
}

// Skip mocks base method
func (m *MockDownloadStreamInterface) Skip(skip int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skip", skip)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Skip indicates an expected call of Skip
func (mr *MockDownloadStreamInterfaceMockRecorder) Skip(skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockDownloadStreamInterface)(nil).Skip), skip)
}

// MockUploadStreamInterface is a mock of UploadStreamInterface interface
type MockUploadStreamInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUploadStreamInterfaceMockRecorder
}

// MockUploadStreamInterfaceMockRecorder is the mock recorder for MockUploadStreamInterface
type MockUploadStreamInterfaceMockRecorder struct {
	mock *MockUploadStreamInterface
}

// NewMockUploadStreamInterface creates a new mock instance
func NewMockUploadStreamInterface(ctrl *gomock.Controller) *MockUploadStreamInterface {
	mock := &MockUploadStreamInterface{ctrl: ctrl}
	mock.recorder = &MockUploadStreamInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUploadStreamInterface) EXPECT() *MockUploadStreamInterfaceMockRecorder {
	return m.recorder
}

// Abort mocks base method
func (m *MockUploadStreamInterface) Abort() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abort")
	ret0, _ := ret[0].(error)
	return ret0
}

// Abort indicates an expected call of Abort
func (mr *MockUploadStreamInterfaceMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockUploadStreamInterface)(nil).Abort))
}

// Close mocks base method
func (m *MockUploadStreamInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUploadStreamInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUploadStreamInterface)(nil).Close))
}

// SetWriteDeadline mocks base method
func (m *MockUploadStreamInterface) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockUploadStreamInterfaceMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockUploadStreamInterface)(nil).SetWriteDeadline), t)
}

// Write mocks base method
func (m *MockUploadStreamInterface) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockUploadStreamInterfaceMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockUploadStreamInterface)(nil).Write), p)
}

// MockBucketInterface is a mock of BucketInterface interface
type MockBucketInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBucketInterfaceMockRecorder
}

// MockBucketInterfaceMockRecorder is the mock recorder for MockBucketInterface
type MockBucketInterfaceMockRecorder struct {
	mock *MockBucketInterface
}

// NewMockBucketInterface creates a new mock instance
func NewMockBucketInterface(ctrl *gomock.Controller) *MockBucketInterface {
	mock := &MockBucketInterface{ctrl: ctrl}
	mock.recorder = &MockBucketInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucketInterface) EXPECT() *MockBucketInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockBucketInterface) Delete(fileID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBucketInterfaceMockRecorder) Delete(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBucketInterface)(nil).Delete), fileID)
}

// DownloadToStream mocks base method
func (m *MockBucketInterface) DownloadToStream(fileID interface{}, stream io.Writer) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadToStream", fileID, stream)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadToStream indicates an expected call of DownloadToStream
func (mr *MockBucketInterfaceMockRecorder) DownloadToStream(fileID, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadToStream", reflect.TypeOf((*MockBucketInterface)(nil).DownloadToStream), fileID, stream)
}

// DownloadToStreamByName mocks base method
func (m *MockBucketInterface) DownloadToStreamByName(filename string, stream io.Writer, opts ...*options.NameOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{filename, stream}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadToStreamByName", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadToStreamByName indicates an expected call of DownloadToStreamByName
func (mr *MockBucketInterfaceMockRecorder) DownloadToStreamByName(filename, stream interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filename, stream}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadToStreamByName", reflect.TypeOf((*MockBucketInterface)(nil).DownloadToStreamByName), varargs...)
}

// Drop mocks base method
func (m *MockBucketInterface) Drop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop
func (mr *MockBucketInterfaceMockRecorder) Drop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockBucketInterface)(nil).Drop))
}

// Find mocks base method
func (m *MockBucketInterface) Find(filter interface{}, opts ...*options.GridFSFindOptions) (pmongo.CursorInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(pmongo.CursorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockBucketInterfaceMockRecorder) Find(filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockBucketInterface)(nil).Find), varargs...)
}

// OpenDownloadStream mocks base method
func (m *MockBucketInterface) OpenDownloadStream(fileID interface{}) (pmongo.DownloadStreamInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenDownloadStream", fileID)
	ret0, _ := ret[0].(pmongo.DownloadStreamInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenDownloadStream indicates an expected call of OpenDownloadStream
func (mr *MockBucketInterfaceMockRecorder) OpenDownloadStream(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDownloadStream", reflect.TypeOf((*MockBucketInterface)(nil).OpenDownloadStream), fileID)
}

// OpenDownloadStreamByName mocks base method
func (m *MockBucketInterface) OpenDownloadStreamByName(filename string, opts ...*options.NameOptions) (pmongo.DownloadStreamInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{filename}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenDownloadStreamByName", varargs...)
	ret0, _ := ret[0].(pmongo.DownloadStreamInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenDownloadStreamByName indicates an expected call of OpenDownloadStreamByName
func (mr *MockBucketInterfaceMockRecorder) OpenDownloadStreamByName(filename interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filename}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDownloadStreamByName", reflect.TypeOf((*MockBucketInterface)(nil).OpenDownloadStreamByName), varargs...)
}

// OpenUploadStream mocks base method
func (m *MockBucketInterface) OpenUploadStream(filename string, opts ...*options.UploadOptions) (pmongo.UploadStreamInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{filename}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenUploadStream", varargs...)
	ret0, _ := ret[0].(pmongo.UploadStreamInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUploadStream indicates an expected call of OpenUploadStream
func (mr *MockBucketInterfaceMockRecorder) OpenUploadStream(filename interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filename}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUploadStream", reflect.TypeOf((*MockBucketInterface)(nil).OpenUploadStream), varargs...)
}

// OpenUploadStreamWithID mocks base method
func (m *MockBucketInterface) OpenUploadStreamWithID(fileID interface{}, filename string, opts ...*options.UploadOptions) (pmongo.UploadStreamInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{fileID, filename}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenUploadStreamWithID", varargs...)
	ret0, _ := ret[0].(pmongo.UploadStreamInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUploadStreamWithID indicates an expected call of OpenUploadStreamWithID
func (mr *MockBucketInterfaceMockRecorder) OpenUploadStreamWithID(fileID, filename interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fileID, filename}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUploadStreamWithID", reflect.TypeOf((*MockBucketInterface)(nil).OpenUploadStreamWithID), varargs...)
}

// Rename mocks base method
func (m *MockBucketInterface) Rename(fileID interface{}, newFilename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", fileID, newFilename)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename
func (mr *MockBucketInterfaceMockRecorder) Rename(fileID, newFilename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockBucketInterface)(nil).Rename), fileID, newFilename)
}

// SetReadDeadline mocks base method
func (m *MockBucketInterface) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockBucketInterfaceMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockBucketInterface)(nil).SetReadDeadline), t)
}

// SetWriteDeadline mocks base method
func (m *MockBucketInterface) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockBucketInterfaceMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockBucketInterface)(nil).SetWriteDeadline), t)
}

// UploadFromStream mocks base method
func (m *MockBucketInterface) UploadFromStream(filename string, source io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{filename, source}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadFromStream", varargs...)
	ret0, _ := ret[0].(primitive.ObjectID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFromStream indicates an expected call of UploadFromStream
func (mr *MockBucketInterfaceMockRecorder) UploadFromStream(filename, source interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filename, source}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFromStream", reflect.TypeOf((*MockBucketInterface)(nil).UploadFromStream), varargs...)
}

// UploadFromStreamWithID mocks base method
func (m *MockBucketInterface) UploadFromStreamWithID(fileID interface{}, filename string, source io.Reader, opts ...*options.UploadOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fileID, filename, source}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadFromStreamWithID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFromStreamWithID indicates an expected call of UploadFromStreamWithID
func (mr *MockBucketInterfaceMockRecorder) UploadFromStreamWithID(fileID, filename, source interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fileID, filename, source}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFromStreamWithID", reflect.TypeOf((*MockBucketInterface)(nil).UploadFromStreamWithID), varargs...)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/connection (interfaces: Connection)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockConnection) Delete(arg0 string, arg1 interface{}) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockConnectionMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockConnection)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockConnection) Get(arg0 string, arg1 connection.APIRequestParameters) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockConnectionMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConnection)(nil).Get), arg0, arg1)
}

// Invoke mocks base method.
func (m *MockConnection) Invoke(arg0 connection.APIRequest) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", arg0)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Invoke indicates an expected call of Invoke.
func (mr *MockConnectionMockRecorder) Invoke(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockConnection)(nil).Invoke), arg0)
}

// Patch mocks base method.
func (m *MockConnection) Patch(arg0 string, arg1 interface{}) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockConnectionMockRecorder) Patch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockConnection)(nil).Patch), arg0, arg1)
}

// Post mocks base method.
func (m *MockConnection) Post(arg0 string, arg1 interface{}) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockConnectionMockRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockConnection)(nil).Post), arg0, arg1)
}

// Put mocks base method.
func (m *MockConnection) Put(arg0 string, arg1 interface{}) (*connection.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(*connection.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockConnectionMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockConnection)(nil).Put), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreis3/users-ms/src/application/operation/user/interface (interfaces: IGetAllUserOperation)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dto "github.com/andreis3/users-ms/src/application/operation/user/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockIGetAllUserOperation is a mock of IGetAllUserOperation interface.
type MockIGetAllUserOperation struct {
	ctrl     *gomock.Controller
	recorder *MockIGetAllUserOperationMockRecorder
}

// MockIGetAllUserOperationMockRecorder is the mock recorder for MockIGetAllUserOperation.
type MockIGetAllUserOperationMockRecorder struct {
	mock *MockIGetAllUserOperation
}

// NewMockIGetAllUserOperation creates a new mock instance.
func NewMockIGetAllUserOperation(ctrl *gomock.Controller) *MockIGetAllUserOperation {
	mock := &MockIGetAllUserOperation{ctrl: ctrl}
	mock.recorder = &MockIGetAllUserOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetAllUserOperation) EXPECT() *MockIGetAllUserOperationMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIGetAllUserOperation) Execute() ([]*dto.UserOutPutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].([]*dto.UserOutPutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIGetAllUserOperationMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIGetAllUserOperation)(nil).Execute))
}
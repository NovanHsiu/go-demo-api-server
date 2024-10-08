// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NovanHsiu/go-demo-api-server/internal/app/service/user (interfaces: UserRepository)

// Package automock is a generated GoMock package.
package automock

import (
	context "context"
	reflect "reflect"

	common "github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	parameter "github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	response "github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUserByParam mocks base method.
func (m *MockUserRepository) CreateUserByParam(arg0 context.Context, arg1 parameter.AddUser) (*response.UserResponseListItem, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserByParam", arg0, arg1)
	ret0, _ := ret[0].(*response.UserResponseListItem)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreateUserByParam indicates an expected call of CreateUserByParam.
func (mr *MockUserRepositoryMockRecorder) CreateUserByParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserByParam", reflect.TypeOf((*MockUserRepository)(nil).CreateUserByParam), arg0, arg1)
}

// DeleteUserByID mocks base method.
func (m *MockUserRepository) DeleteUserByID(arg0 context.Context, arg1 string) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", arg0, arg1)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// DeleteUserByID indicates an expected call of DeleteUserByID.
func (mr *MockUserRepositoryMockRecorder) DeleteUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockUserRepository)(nil).DeleteUserByID), arg0, arg1)
}

// GetUserAndPasswordByAccount mocks base method.
func (m *MockUserRepository) GetUserAndPasswordByAccount(arg0 context.Context, arg1 string) (*response.UserResponseListItem, string, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAndPasswordByAccount", arg0, arg1)
	ret0, _ := ret[0].(*response.UserResponseListItem)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(common.Error)
	return ret0, ret1, ret2
}

// GetUserAndPasswordByAccount indicates an expected call of GetUserAndPasswordByAccount.
func (mr *MockUserRepositoryMockRecorder) GetUserAndPasswordByAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAndPasswordByAccount", reflect.TypeOf((*MockUserRepository)(nil).GetUserAndPasswordByAccount), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(arg0 context.Context, arg1 string) (*response.UserResponseListItem, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*response.UserResponseListItem)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), arg0, arg1)
}

// GetUsersByParam mocks base method.
func (m *MockUserRepository) GetUsersByParam(arg0 context.Context, arg1 parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByParam", arg0, arg1)
	ret0, _ := ret[0].([]response.UserResponseListItem)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(common.Error)
	return ret0, ret1, ret2
}

// GetUsersByParam indicates an expected call of GetUsersByParam.
func (mr *MockUserRepositoryMockRecorder) GetUsersByParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByParam", reflect.TypeOf((*MockUserRepository)(nil).GetUsersByParam), arg0, arg1)
}

// UpdateUserByIDnParam mocks base method.
func (m *MockUserRepository) UpdateUserByIDnParam(arg0 context.Context, arg1 string, arg2 parameter.ModifyUser) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByIDnParam", arg0, arg1, arg2)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// UpdateUserByIDnParam indicates an expected call of UpdateUserByIDnParam.
func (mr *MockUserRepositoryMockRecorder) UpdateUserByIDnParam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByIDnParam", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserByIDnParam), arg0, arg1, arg2)
}

// UpdateUserPasswordByIDnParam mocks base method.
func (m *MockUserRepository) UpdateUserPasswordByIDnParam(arg0 context.Context, arg1 string, arg2 parameter.ModifyUserPassword) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPasswordByIDnParam", arg0, arg1, arg2)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// UpdateUserPasswordByIDnParam indicates an expected call of UpdateUserPasswordByIDnParam.
func (mr *MockUserRepositoryMockRecorder) UpdateUserPasswordByIDnParam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPasswordByIDnParam", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserPasswordByIDnParam), arg0, arg1, arg2)
}

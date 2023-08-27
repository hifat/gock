// Code generated by MockGen. DO NOT EDIT.
// Source: ./task.go

// Package mockTaskService is a generated GoMock package.
package mockTaskService

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	taskDomain "github.com/hifat/gock/internal/domain/taskDomain"
	gomock "go.uber.org/mock/gomock"
)

// MockITaskService is a mock of ITaskService interface.
type MockITaskService struct {
	ctrl     *gomock.Controller
	recorder *MockITaskServiceMockRecorder
}

// MockITaskServiceMockRecorder is the mock recorder for MockITaskService.
type MockITaskServiceMockRecorder struct {
	mock *MockITaskService
}

// NewMockITaskService creates a new mock instance.
func NewMockITaskService(ctrl *gomock.Controller) *MockITaskService {
	mock := &MockITaskService{ctrl: ctrl}
	mock.recorder = &MockITaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskService) EXPECT() *MockITaskServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITaskService) Create(ctx context.Context, req *taskDomain.TaskRequest) (*taskDomain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(*taskDomain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITaskServiceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITaskService)(nil).Create), ctx, req)
}

// Delete mocks base method.
func (m *MockITaskService) Delete(ctx context.Context, taskID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITaskServiceMockRecorder) Delete(ctx, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITaskService)(nil).Delete), ctx, taskID)
}

// Get mocks base method.
func (m *MockITaskService) Get(ctx context.Context, res *[]taskDomain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, res)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockITaskServiceMockRecorder) Get(ctx, res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockITaskService)(nil).Get), ctx, res)
}

// GetByID mocks base method.
func (m *MockITaskService) GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, res, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockITaskServiceMockRecorder) GetByID(ctx, res, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockITaskService)(nil).GetByID), ctx, res, taskID)
}

// Update mocks base method.
func (m *MockITaskService) Update(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req, taskID)
	ret0, _ := ret[0].(*taskDomain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockITaskServiceMockRecorder) Update(ctx, req, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITaskService)(nil).Update), ctx, req, taskID)
}

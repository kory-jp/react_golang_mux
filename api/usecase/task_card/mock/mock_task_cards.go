// Code generated by MockGen. DO NOT EDIT.
// Source: task_card_repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/kory-jp/react_golang_mux/api/domain"
)

// MockTaskCardRepository is a mock of TaskCardRepository interface.
type MockTaskCardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskCardRepositoryMockRecorder
}

// MockTaskCardRepositoryMockRecorder is the mock recorder for MockTaskCardRepository.
type MockTaskCardRepositoryMockRecorder struct {
	mock *MockTaskCardRepository
}

// NewMockTaskCardRepository creates a new mock instance.
func NewMockTaskCardRepository(ctrl *gomock.Controller) *MockTaskCardRepository {
	mock := &MockTaskCardRepository{ctrl: ctrl}
	mock.recorder = &MockTaskCardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskCardRepository) EXPECT() *MockTaskCardRepositoryMockRecorder {
	return m.recorder
}

// ChangeBoolean mocks base method.
func (m *MockTaskCardRepository) ChangeBoolean(arg0, arg1 int, arg2 domain.TaskCard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeBoolean", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeBoolean indicates an expected call of ChangeBoolean.
func (mr *MockTaskCardRepositoryMockRecorder) ChangeBoolean(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeBoolean", reflect.TypeOf((*MockTaskCardRepository)(nil).ChangeBoolean), arg0, arg1, arg2)
}

// Erasure mocks base method.
func (m *MockTaskCardRepository) Erasure(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Erasure", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Erasure indicates an expected call of Erasure.
func (mr *MockTaskCardRepositoryMockRecorder) Erasure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Erasure", reflect.TypeOf((*MockTaskCardRepository)(nil).Erasure), arg0, arg1)
}

// FindByIdAndUserId mocks base method.
func (m *MockTaskCardRepository) FindByIdAndUserId(arg0, arg1 int) (*domain.TaskCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdAndUserId", arg0, arg1)
	ret0, _ := ret[0].(*domain.TaskCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdAndUserId indicates an expected call of FindByIdAndUserId.
func (mr *MockTaskCardRepositoryMockRecorder) FindByIdAndUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdAndUserId", reflect.TypeOf((*MockTaskCardRepository)(nil).FindByIdAndUserId), arg0, arg1)
}

// FindByTodoIdAndUserId mocks base method.
func (m *MockTaskCardRepository) FindByTodoIdAndUserId(arg0, arg1, arg2 int) (domain.TaskCards, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTodoIdAndUserId", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.TaskCards)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindByTodoIdAndUserId indicates an expected call of FindByTodoIdAndUserId.
func (mr *MockTaskCardRepositoryMockRecorder) FindByTodoIdAndUserId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTodoIdAndUserId", reflect.TypeOf((*MockTaskCardRepository)(nil).FindByTodoIdAndUserId), arg0, arg1, arg2)
}

// Overwrite mocks base method.
func (m *MockTaskCardRepository) Overwrite(arg0 domain.TaskCard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Overwrite", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Overwrite indicates an expected call of Overwrite.
func (mr *MockTaskCardRepositoryMockRecorder) Overwrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overwrite", reflect.TypeOf((*MockTaskCardRepository)(nil).Overwrite), arg0)
}

// Store mocks base method.
func (m *MockTaskCardRepository) Store(arg0 domain.TaskCard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTaskCardRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTaskCardRepository)(nil).Store), arg0)
}
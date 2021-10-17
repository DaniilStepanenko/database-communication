// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository (interfaces: CustomerRepository)

// Package mockgen is a generated GoMock package.
package mockgen

import (
	context "context"
	reflect "reflect"
	time "time"

	repository "github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	model "github.com/DaniilStepanenko/database-communication/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerRepository) Create(arg0 context.Context, arg1 *model.Customer) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), arg0, arg1)
}

// FindActiveCustomers mocks base method.
func (m *MockCustomerRepository) FindActiveCustomers(arg0 context.Context, arg1 time.Time) ([]*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActiveCustomers", arg0, arg1)
	ret0, _ := ret[0].([]*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActiveCustomers indicates an expected call of FindActiveCustomers.
func (mr *MockCustomerRepositoryMockRecorder) FindActiveCustomers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActiveCustomers", reflect.TypeOf((*MockCustomerRepository)(nil).FindActiveCustomers), arg0, arg1)
}

// List mocks base method.
func (m *MockCustomerRepository) List(arg0 context.Context, arg1 *repository.ListOptions, arg2 *repository.CustomerCriteria) ([]*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCustomerRepositoryMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomerRepository)(nil).List), arg0, arg1, arg2)
}

// Read mocks base method.
func (m *MockCustomerRepository) Read(arg0 context.Context, arg1 int) (*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCustomerRepositoryMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCustomerRepository)(nil).Read), arg0, arg1)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(arg0 context.Context, arg1 *model.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), arg0, arg1)
}

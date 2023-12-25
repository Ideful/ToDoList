// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "ToDoList/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockToDoList is a mock of ToDoList interface.
type MockToDoList struct {
	ctrl     *gomock.Controller
	recorder *MockToDoListMockRecorder
}

// MockToDoListMockRecorder is the mock recorder for MockToDoList.
type MockToDoListMockRecorder struct {
	mock *MockToDoList
}

// NewMockToDoList creates a new mock instance.
func NewMockToDoList(ctrl *gomock.Controller) *MockToDoList {
	mock := &MockToDoList{ctrl: ctrl}
	mock.recorder = &MockToDoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToDoList) EXPECT() *MockToDoListMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockToDoList) Create(userId int, list model.TodoList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, list)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockToDoListMockRecorder) Create(userId, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockToDoList)(nil).Create), userId, list)
}

// Delete mocks base method.
func (m *MockToDoList) Delete(userId, listId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, listId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockToDoListMockRecorder) Delete(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockToDoList)(nil).Delete), userId, listId)
}

// GetAll mocks base method.
func (m *MockToDoList) GetAll(userId int) ([]model.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]model.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockToDoListMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockToDoList)(nil).GetAll), userId)
}

// GetById mocks base method.
func (m *MockToDoList) GetById(userId, listId int) (model.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, listId)
	ret0, _ := ret[0].(model.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockToDoListMockRecorder) GetById(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockToDoList)(nil).GetById), userId, listId)
}

// Update mocks base method.
func (m *MockToDoList) Update(userId, id int, input model.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockToDoListMockRecorder) Update(userId, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockToDoList)(nil).Update), userId, id, input)
}

// MockToDoItem is a mock of ToDoItem interface.
type MockToDoItem struct {
	ctrl     *gomock.Controller
	recorder *MockToDoItemMockRecorder
}

// MockToDoItemMockRecorder is the mock recorder for MockToDoItem.
type MockToDoItemMockRecorder struct {
	mock *MockToDoItem
}

// NewMockToDoItem creates a new mock instance.
func NewMockToDoItem(ctrl *gomock.Controller) *MockToDoItem {
	mock := &MockToDoItem{ctrl: ctrl}
	mock.recorder = &MockToDoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToDoItem) EXPECT() *MockToDoItemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockToDoItem) Create(userId, itemId int, item model.TodoItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, itemId, item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockToDoItemMockRecorder) Create(userId, itemId, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockToDoItem)(nil).Create), userId, itemId, item)
}

// Delete mocks base method.
func (m *MockToDoItem) Delete(userId, itemId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockToDoItemMockRecorder) Delete(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockToDoItem)(nil).Delete), userId, itemId)
}

// GetAll mocks base method.
func (m *MockToDoItem) GetAll(userId, itemId int) ([]model.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId, itemId)
	ret0, _ := ret[0].([]model.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockToDoItemMockRecorder) GetAll(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockToDoItem)(nil).GetAll), userId, itemId)
}

// GetById mocks base method.
func (m *MockToDoItem) GetById(userId, itemId int) (model.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, itemId)
	ret0, _ := ret[0].(model.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockToDoItemMockRecorder) GetById(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockToDoItem)(nil).GetById), userId, itemId)
}

// Update mocks base method.
func (m *MockToDoItem) Update(userId, itemId int, input model.UpdateItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, itemId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockToDoItemMockRecorder) Update(userId, itemId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockToDoItem)(nil).Update), userId, itemId, input)
}
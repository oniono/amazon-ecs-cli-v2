// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/archer/env.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEnvironmentStore is a mock of EnvironmentStore interface
type MockEnvironmentStore struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentStoreMockRecorder
}

// MockEnvironmentStoreMockRecorder is the mock recorder for MockEnvironmentStore
type MockEnvironmentStoreMockRecorder struct {
	mock *MockEnvironmentStore
}

// NewMockEnvironmentStore creates a new mock instance
func NewMockEnvironmentStore(ctrl *gomock.Controller) *MockEnvironmentStore {
	mock := &MockEnvironmentStore{ctrl: ctrl}
	mock.recorder = &MockEnvironmentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentStore) EXPECT() *MockEnvironmentStoreMockRecorder {
	return m.recorder
}

// ListEnvironments mocks base method
func (m *MockEnvironmentStore) ListEnvironments(projectName string) ([]*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnvironments", projectName)
	ret0, _ := ret[0].([]*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments
func (mr *MockEnvironmentStoreMockRecorder) ListEnvironments(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockEnvironmentStore)(nil).ListEnvironments), projectName)
}

// GetEnvironment mocks base method
func (m *MockEnvironmentStore) GetEnvironment(projectName, environmentName string) (*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockEnvironmentStoreMockRecorder) GetEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentStore)(nil).GetEnvironment), projectName, environmentName)
}

// CreateEnvironment mocks base method
func (m *MockEnvironmentStore) CreateEnvironment(env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvironment indicates an expected call of CreateEnvironment
func (mr *MockEnvironmentStoreMockRecorder) CreateEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentStore)(nil).CreateEnvironment), env)
}

// DeleteEnvironment mocks base method
func (m *MockEnvironmentStore) DeleteEnvironment(projectName, environmentName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment
func (mr *MockEnvironmentStoreMockRecorder) DeleteEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockEnvironmentStore)(nil).DeleteEnvironment), projectName, environmentName)
}

// MockEnvironmentLister is a mock of EnvironmentLister interface
type MockEnvironmentLister struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentListerMockRecorder
}

// MockEnvironmentListerMockRecorder is the mock recorder for MockEnvironmentLister
type MockEnvironmentListerMockRecorder struct {
	mock *MockEnvironmentLister
}

// NewMockEnvironmentLister creates a new mock instance
func NewMockEnvironmentLister(ctrl *gomock.Controller) *MockEnvironmentLister {
	mock := &MockEnvironmentLister{ctrl: ctrl}
	mock.recorder = &MockEnvironmentListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentLister) EXPECT() *MockEnvironmentListerMockRecorder {
	return m.recorder
}

// ListEnvironments mocks base method
func (m *MockEnvironmentLister) ListEnvironments(projectName string) ([]*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnvironments", projectName)
	ret0, _ := ret[0].([]*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments
func (mr *MockEnvironmentListerMockRecorder) ListEnvironments(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockEnvironmentLister)(nil).ListEnvironments), projectName)
}

// MockEnvironmentGetter is a mock of EnvironmentGetter interface
type MockEnvironmentGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentGetterMockRecorder
}

// MockEnvironmentGetterMockRecorder is the mock recorder for MockEnvironmentGetter
type MockEnvironmentGetterMockRecorder struct {
	mock *MockEnvironmentGetter
}

// NewMockEnvironmentGetter creates a new mock instance
func NewMockEnvironmentGetter(ctrl *gomock.Controller) *MockEnvironmentGetter {
	mock := &MockEnvironmentGetter{ctrl: ctrl}
	mock.recorder = &MockEnvironmentGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentGetter) EXPECT() *MockEnvironmentGetterMockRecorder {
	return m.recorder
}

// GetEnvironment mocks base method
func (m *MockEnvironmentGetter) GetEnvironment(projectName, environmentName string) (*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockEnvironmentGetterMockRecorder) GetEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentGetter)(nil).GetEnvironment), projectName, environmentName)
}

// MockEnvironmentCreator is a mock of EnvironmentCreator interface
type MockEnvironmentCreator struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentCreatorMockRecorder
}

// MockEnvironmentCreatorMockRecorder is the mock recorder for MockEnvironmentCreator
type MockEnvironmentCreatorMockRecorder struct {
	mock *MockEnvironmentCreator
}

// NewMockEnvironmentCreator creates a new mock instance
func NewMockEnvironmentCreator(ctrl *gomock.Controller) *MockEnvironmentCreator {
	mock := &MockEnvironmentCreator{ctrl: ctrl}
	mock.recorder = &MockEnvironmentCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentCreator) EXPECT() *MockEnvironmentCreatorMockRecorder {
	return m.recorder
}

// CreateEnvironment mocks base method
func (m *MockEnvironmentCreator) CreateEnvironment(env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvironment indicates an expected call of CreateEnvironment
func (mr *MockEnvironmentCreatorMockRecorder) CreateEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentCreator)(nil).CreateEnvironment), env)
}

// MockEnvironmentDeleter is a mock of EnvironmentDeleter interface
type MockEnvironmentDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentDeleterMockRecorder
}

// MockEnvironmentDeleterMockRecorder is the mock recorder for MockEnvironmentDeleter
type MockEnvironmentDeleterMockRecorder struct {
	mock *MockEnvironmentDeleter
}

// NewMockEnvironmentDeleter creates a new mock instance
func NewMockEnvironmentDeleter(ctrl *gomock.Controller) *MockEnvironmentDeleter {
	mock := &MockEnvironmentDeleter{ctrl: ctrl}
	mock.recorder = &MockEnvironmentDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentDeleter) EXPECT() *MockEnvironmentDeleterMockRecorder {
	return m.recorder
}

// DeleteEnvironment mocks base method
func (m *MockEnvironmentDeleter) DeleteEnvironment(projectName, environmentName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment
func (mr *MockEnvironmentDeleterMockRecorder) DeleteEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockEnvironmentDeleter)(nil).DeleteEnvironment), projectName, environmentName)
}

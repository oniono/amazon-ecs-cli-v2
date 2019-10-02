// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/term/spinner/spinner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockspinner is a mock of spinner interface
type Mockspinner struct {
	ctrl     *gomock.Controller
	recorder *MockspinnerMockRecorder
}

// MockspinnerMockRecorder is the mock recorder for Mockspinner
type MockspinnerMockRecorder struct {
	mock *Mockspinner
}

// NewMockspinner creates a new mock instance
func NewMockspinner(ctrl *gomock.Controller) *Mockspinner {
	mock := &Mockspinner{ctrl: ctrl}
	mock.recorder = &MockspinnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockspinner) EXPECT() *MockspinnerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *Mockspinner) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockspinnerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockspinner)(nil).Start))
}

// Stop mocks base method
func (m *Mockspinner) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockspinnerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockspinner)(nil).Stop))
}

// MockwriteFlusher is a mock of writeFlusher interface
type MockwriteFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockwriteFlusherMockRecorder
}

// MockwriteFlusherMockRecorder is the mock recorder for MockwriteFlusher
type MockwriteFlusherMockRecorder struct {
	mock *MockwriteFlusher
}

// NewMockwriteFlusher creates a new mock instance
func NewMockwriteFlusher(ctrl *gomock.Controller) *MockwriteFlusher {
	mock := &MockwriteFlusher{ctrl: ctrl}
	mock.recorder = &MockwriteFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockwriteFlusher) EXPECT() *MockwriteFlusherMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockwriteFlusher) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockwriteFlusherMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockwriteFlusher)(nil).Write), p)
}

// Flush mocks base method
func (m *MockwriteFlusher) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockwriteFlusherMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockwriteFlusher)(nil).Flush))
}
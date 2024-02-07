// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/http/v2 (interfaces: HTTPClient,RequestRecorder,Logger)
//
// Generated by this command:
//
//	mockgen -package http -destination client_mock_test.go . HTTPClient,RequestRecorder,Logger
//

// Package http is a generated GoMock package.
package http

import (
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPClient) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), arg0)
}

// MockRequestRecorder is a mock of RequestRecorder interface.
type MockRequestRecorder struct {
	ctrl     *gomock.Controller
	recorder *MockRequestRecorderMockRecorder
}

// MockRequestRecorderMockRecorder is the mock recorder for MockRequestRecorder.
type MockRequestRecorderMockRecorder struct {
	mock *MockRequestRecorder
}

// NewMockRequestRecorder creates a new mock instance.
func NewMockRequestRecorder(ctrl *gomock.Controller) *MockRequestRecorder {
	mock := &MockRequestRecorder{ctrl: ctrl}
	mock.recorder = &MockRequestRecorderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestRecorder) EXPECT() *MockRequestRecorderMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockRequestRecorder) Record(arg0 string, arg1 *url.URL, arg2 *http.Response, arg3 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Record", arg0, arg1, arg2, arg3)
}

// Record indicates an expected call of Record.
func (mr *MockRequestRecorderMockRecorder) Record(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockRequestRecorder)(nil).Record), arg0, arg1, arg2, arg3)
}

// RecordError mocks base method.
func (m *MockRequestRecorder) RecordError(arg0 string, arg1 *url.URL, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordError", arg0, arg1, arg2)
}

// RecordError indicates an expected call of RecordError.
func (mr *MockRequestRecorderMockRecorder) RecordError(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordError", reflect.TypeOf((*MockRequestRecorder)(nil).RecordError), arg0, arg1, arg2)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}

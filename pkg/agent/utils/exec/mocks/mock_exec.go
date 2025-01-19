/*
Copyright 2024 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: exec.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExecInterface is a mock of ExecInterface interface.
type MockExecInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExecInterfaceMockRecorder
}

// MockExecInterfaceMockRecorder is the mock recorder for MockExecInterface.
type MockExecInterfaceMockRecorder struct {
	mock *MockExecInterface
}

// NewMockExecInterface creates a new mock instance.
func NewMockExecInterface(ctrl *gomock.Controller) *MockExecInterface {
	mock := &MockExecInterface{ctrl: ctrl}
	mock.recorder = &MockExecInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecInterface) EXPECT() *MockExecInterfaceMockRecorder {
	return m.recorder
}

// CommandContext mocks base method.
func (m *MockExecInterface) CommandContext(ctx context.Context, cmd string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandContext", ctx, cmd)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandContext indicates an expected call of CommandContext.
func (mr *MockExecInterfaceMockRecorder) CommandContext(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockExecInterface)(nil).CommandContext), ctx, cmd)
}

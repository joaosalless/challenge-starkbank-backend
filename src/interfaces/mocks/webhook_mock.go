// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/interfaces/webhook.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	dtos "github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

// MockWebhookHandler is a mock of WebhookHandler interface.
type MockWebhookHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookHandlerMockRecorder
}

// MockWebhookHandlerMockRecorder is the mock recorder for MockWebhookHandler.
type MockWebhookHandlerMockRecorder struct {
	mock *MockWebhookHandler
}

// NewMockWebhookHandler creates a new mock instance.
func NewMockWebhookHandler(ctrl *gomock.Controller) *MockWebhookHandler {
	mock := &MockWebhookHandler{ctrl: ctrl}
	mock.recorder = &MockWebhookHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookHandler) EXPECT() *MockWebhookHandlerMockRecorder {
	return m.recorder
}

// ProcessEvent mocks base method.
func (m *MockWebhookHandler) ProcessEvent(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessEvent", ctx)
}

// ProcessEvent indicates an expected call of ProcessEvent.
func (mr *MockWebhookHandlerMockRecorder) ProcessEvent(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessEvent", reflect.TypeOf((*MockWebhookHandler)(nil).ProcessEvent), ctx)
}

// MockWebhookController is a mock of WebhookController interface.
type MockWebhookController struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookControllerMockRecorder
}

// MockWebhookControllerMockRecorder is the mock recorder for MockWebhookController.
type MockWebhookControllerMockRecorder struct {
	mock *MockWebhookController
}

// NewMockWebhookController creates a new mock instance.
func NewMockWebhookController(ctrl *gomock.Controller) *MockWebhookController {
	mock := &MockWebhookController{ctrl: ctrl}
	mock.recorder = &MockWebhookControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookController) EXPECT() *MockWebhookControllerMockRecorder {
	return m.recorder
}

// ProcessEvent mocks base method.
func (m *MockWebhookController) ProcessEvent(ctx context.Context, input dtos.WebhookProcessEventInput) (dtos.WebhookProcessEventOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessEvent", ctx, input)
	ret0, _ := ret[0].(dtos.WebhookProcessEventOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessEvent indicates an expected call of ProcessEvent.
func (mr *MockWebhookControllerMockRecorder) ProcessEvent(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessEvent", reflect.TypeOf((*MockWebhookController)(nil).ProcessEvent), ctx, input)
}

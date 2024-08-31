// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/interfaces/invoice.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	dtos "joaosalless/challenge-starkbank/src/dtos"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockInvoiceHandler is a mock of InvoiceHandler interface.
type MockInvoiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceHandlerMockRecorder
}

// MockInvoiceHandlerMockRecorder is the mock recorder for MockInvoiceHandler.
type MockInvoiceHandlerMockRecorder struct {
	mock *MockInvoiceHandler
}

// NewMockInvoiceHandler creates a new mock instance.
func NewMockInvoiceHandler(ctrl *gomock.Controller) *MockInvoiceHandler {
	mock := &MockInvoiceHandler{ctrl: ctrl}
	mock.recorder = &MockInvoiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceHandler) EXPECT() *MockInvoiceHandlerMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method.
func (m *MockInvoiceHandler) CreateInvoice(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateInvoice", ctx)
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockInvoiceHandlerMockRecorder) CreateInvoice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceHandler)(nil).CreateInvoice), ctx)
}

// HookProcessInvoice mocks base method.
func (m *MockInvoiceHandler) HookProcessInvoice(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HookProcessInvoice", ctx)
}

// HookProcessInvoice indicates an expected call of HookProcessInvoice.
func (mr *MockInvoiceHandlerMockRecorder) HookProcessInvoice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HookProcessInvoice", reflect.TypeOf((*MockInvoiceHandler)(nil).HookProcessInvoice), ctx)
}

// MockInvoiceController is a mock of InvoiceController interface.
type MockInvoiceController struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceControllerMockRecorder
}

// MockInvoiceControllerMockRecorder is the mock recorder for MockInvoiceController.
type MockInvoiceControllerMockRecorder struct {
	mock *MockInvoiceController
}

// NewMockInvoiceController creates a new mock instance.
func NewMockInvoiceController(ctrl *gomock.Controller) *MockInvoiceController {
	mock := &MockInvoiceController{ctrl: ctrl}
	mock.recorder = &MockInvoiceControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceController) EXPECT() *MockInvoiceControllerMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method.
func (m *MockInvoiceController) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (dtos.CreateInvoiceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", ctx, input)
	ret0, _ := ret[0].(dtos.CreateInvoiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockInvoiceControllerMockRecorder) CreateInvoice(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceController)(nil).CreateInvoice), ctx, input)
}

// InvoiceHookProcess mocks base method.
func (m *MockInvoiceController) InvoiceHookProcess(ctx context.Context, input dtos.InvoiceHookProcessInput) (dtos.InvoiceHookProcessOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvoiceHookProcess", ctx, input)
	ret0, _ := ret[0].(dtos.InvoiceHookProcessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvoiceHookProcess indicates an expected call of InvoiceHookProcess.
func (mr *MockInvoiceControllerMockRecorder) InvoiceHookProcess(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceHookProcess", reflect.TypeOf((*MockInvoiceController)(nil).InvoiceHookProcess), ctx, input)
}

// MockInvoiceService is a mock of InvoiceService interface.
type MockInvoiceService struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceServiceMockRecorder
}

// MockInvoiceServiceMockRecorder is the mock recorder for MockInvoiceService.
type MockInvoiceServiceMockRecorder struct {
	mock *MockInvoiceService
}

// NewMockInvoiceService creates a new mock instance.
func NewMockInvoiceService(ctrl *gomock.Controller) *MockInvoiceService {
	mock := &MockInvoiceService{ctrl: ctrl}
	mock.recorder = &MockInvoiceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceService) EXPECT() *MockInvoiceServiceMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method.
func (m *MockInvoiceService) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (dtos.CreateInvoiceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", ctx, input)
	ret0, _ := ret[0].(dtos.CreateInvoiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockInvoiceServiceMockRecorder) CreateInvoice(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceService)(nil).CreateInvoice), ctx, input)
}

// MockInvoiceCreateCron is a mock of InvoiceCreateCron interface.
type MockInvoiceCreateCron struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceCreateCronMockRecorder
}

// MockInvoiceCreateCronMockRecorder is the mock recorder for MockInvoiceCreateCron.
type MockInvoiceCreateCronMockRecorder struct {
	mock *MockInvoiceCreateCron
}

// NewMockInvoiceCreateCron creates a new mock instance.
func NewMockInvoiceCreateCron(ctrl *gomock.Controller) *MockInvoiceCreateCron {
	mock := &MockInvoiceCreateCron{ctrl: ctrl}
	mock.recorder = &MockInvoiceCreateCronMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceCreateCron) EXPECT() *MockInvoiceCreateCronMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockInvoiceCreateCron) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockInvoiceCreateCronMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockInvoiceCreateCron)(nil).Run))
}

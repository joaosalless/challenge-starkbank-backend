// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/interfaces/bank_gateway.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dtos "github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	interfaces "github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	privatekey "github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

// MockBankGateway is a mock of BankGateway interface.
type MockBankGateway struct {
	ctrl     *gomock.Controller
	recorder *MockBankGatewayMockRecorder
}

// MockBankGatewayMockRecorder is the mock recorder for MockBankGateway.
type MockBankGatewayMockRecorder struct {
	mock *MockBankGateway
}

// NewMockBankGateway creates a new mock instance.
func NewMockBankGateway(ctrl *gomock.Controller) *MockBankGateway {
	mock := &MockBankGateway{ctrl: ctrl}
	mock.recorder = &MockBankGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankGateway) EXPECT() *MockBankGatewayMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method.
func (m *MockBankGateway) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (dtos.CreateInvoiceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", ctx, input)
	ret0, _ := ret[0].(dtos.CreateInvoiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockBankGatewayMockRecorder) CreateInvoice(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockBankGateway)(nil).CreateInvoice), ctx, input)
}

// CreateTransfer mocks base method.
func (m *MockBankGateway) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (dtos.CreateTransferOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", ctx, input)
	ret0, _ := ret[0].(dtos.CreateTransferOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockBankGatewayMockRecorder) CreateTransfer(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockBankGateway)(nil).CreateTransfer), ctx, input)
}

// GetUser mocks base method.
func (m *MockBankGateway) GetUser() interfaces.BankGatewayUser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].(interfaces.BankGatewayUser)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockBankGatewayMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockBankGateway)(nil).GetUser))
}

// MockBankGatewayUser is a mock of BankGatewayUser interface.
type MockBankGatewayUser struct {
	ctrl     *gomock.Controller
	recorder *MockBankGatewayUserMockRecorder
}

// MockBankGatewayUserMockRecorder is the mock recorder for MockBankGatewayUser.
type MockBankGatewayUserMockRecorder struct {
	mock *MockBankGatewayUser
}

// NewMockBankGatewayUser creates a new mock instance.
func NewMockBankGatewayUser(ctrl *gomock.Controller) *MockBankGatewayUser {
	mock := &MockBankGatewayUser{ctrl: ctrl}
	mock.recorder = &MockBankGatewayUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankGatewayUser) EXPECT() *MockBankGatewayUserMockRecorder {
	return m.recorder
}

// GetAcessId mocks base method.
func (m *MockBankGatewayUser) GetAcessId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcessId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAcessId indicates an expected call of GetAcessId.
func (mr *MockBankGatewayUserMockRecorder) GetAcessId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcessId", reflect.TypeOf((*MockBankGatewayUser)(nil).GetAcessId))
}

// GetEnvironment mocks base method.
func (m *MockBankGatewayUser) GetEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockBankGatewayUserMockRecorder) GetEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockBankGatewayUser)(nil).GetEnvironment))
}

// GetPrivateKey mocks base method.
func (m *MockBankGatewayUser) GetPrivateKey() *privatekey.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateKey")
	ret0, _ := ret[0].(*privatekey.PrivateKey)
	return ret0
}

// GetPrivateKey indicates an expected call of GetPrivateKey.
func (mr *MockBankGatewayUserMockRecorder) GetPrivateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateKey", reflect.TypeOf((*MockBankGatewayUser)(nil).GetPrivateKey))
}

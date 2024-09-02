package services

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type InvoiceService struct {
	app         interfaces.Application
	bankGateway interfaces.BankGateway
}

type InvoiceServiceDependencies struct {
	ioc.In
	Application interfaces.Application `name:"Application"`
	BankGateway interfaces.BankGateway `name:"BankGateway"`
}

func NewInvoiceService(deps InvoiceServiceDependencies) *InvoiceService {
	return &InvoiceService{app: deps.Application, bankGateway: deps.BankGateway}
}

func (is InvoiceService) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	is.app.Logger().Infow("InvoiceService.CreateInvoice", "input", input)
	return is.bankGateway.CreateInvoice(ctx, input)
}

package services

import (
	"context"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceService struct {
	logger      interfaces.Logger
	bankGateway interfaces.BankGateway
}

type InvoiceServiceDependencies struct {
	app.Dependencies
	BankGateway interfaces.BankGateway `name:"BankGateway"`
}

func NewInvoiceService(deps InvoiceServiceDependencies) *InvoiceService {
	return &InvoiceService{logger: deps.Logger, bankGateway: deps.BankGateway}
}

func (is InvoiceService) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	is.logger.Infow("InvoiceService.CreateInvoice", "input", input)
	return is.bankGateway.CreateInvoice(ctx, input)
}

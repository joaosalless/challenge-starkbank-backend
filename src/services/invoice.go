package services

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/app"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
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

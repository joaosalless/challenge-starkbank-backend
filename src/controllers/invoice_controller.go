package controllers

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/application"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type InvoiceController struct {
	logger          interfaces.Logger
	invoiceService  interfaces.InvoiceService
	transferService interfaces.TransferService
}

type InvoiceControllerDependencies struct {
	application.Dependencies
	InvoiceService  interfaces.InvoiceService  `name:"InvoiceService"`
	TransferService interfaces.TransferService `name:"TransferService"`
}

func NewInvoiceController(deps InvoiceControllerDependencies) *InvoiceController {
	return &InvoiceController{
		logger:          deps.Logger,
		invoiceService:  deps.InvoiceService,
		transferService: deps.TransferService,
	}
}

func (i *InvoiceController) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	i.logger.Infow("InvoiceController.CreateInvoice", "input", input)
	return i.invoiceService.CreateInvoice(ctx, input)
}

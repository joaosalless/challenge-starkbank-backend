package controllers

import (
	"context"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceController struct {
	logger         interfaces.Logger
	invoiceService interfaces.InvoiceService
}

type InvoiceControllerDependencies struct {
	app.Dependencies
	InvoiceService interfaces.InvoiceService `name:"InvoiceService"`
}

func NewInvoiceController(deps InvoiceControllerDependencies) *InvoiceController {
	return &InvoiceController{logger: deps.Logger, invoiceService: deps.InvoiceService}
}

func (i *InvoiceController) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	i.logger.Infow("InvoiceController.CreateInvoice", "input", input)
	return i.invoiceService.CreateInvoice(ctx, input)
}

func (i *InvoiceController) InvoiceHookProcess(ctx context.Context, input dtos.InvoiceHookProcessInput) (output dtos.InvoiceHookProcessOutput, err error) {
	i.logger.Infow("InvoiceController.InvoiceHookProcess", "input", input)
	return
}

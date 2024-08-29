package controllers

import (
	"context"
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceController struct {
	invoiceService  interfaces.InvoiceService
	transferService interfaces.TransferService
}

type InvoiceControllerDependencies struct {
	dig.In
	InvoiceService  interfaces.InvoiceService  `name:"InvoiceService"`
	TransferService interfaces.TransferService `name:"TransferService"`
}

func NewInvoiceController(deps InvoiceControllerDependencies) *InvoiceController {
	return &InvoiceController{invoiceService: deps.InvoiceService, transferService: deps.TransferService}
}

func (i *InvoiceController) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	return
}

func (i *InvoiceController) InvoiceHookProcess(ctx context.Context, input dtos.InvoiceHookProcessInput) (output dtos.InvoiceHookProcessOutput, err error) {
	return
}

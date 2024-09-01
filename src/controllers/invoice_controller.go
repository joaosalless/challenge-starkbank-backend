package controllers

import (
	"context"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/domain"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceController struct {
	logger          interfaces.Logger
	invoiceService  interfaces.InvoiceService
	transferService interfaces.TransferService
}

type InvoiceControllerDependencies struct {
	app.Dependencies
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

func (i *InvoiceController) InvoiceHookProcess(ctx context.Context, input dtos.InvoiceHookProcessInput) (output dtos.InvoiceHookProcessOutput, err error) {
	i.logger.Infow("InvoiceController.InvoiceHookProcess", "input", input)

	if input.Event.Type == domain.InvoiceEventCredited {
		transferInput := dtos.CreateTransferFromInvoiceInput{Data: input.Event.Log.Invoice}

		_, err = i.transferService.CreateTransferFromInvoice(ctx, transferInput)
		if err != nil {
			i.logger.Errorw("Failed to CreateTransferFromInvoice", "error", err)
			return
		}
	}

	return
}

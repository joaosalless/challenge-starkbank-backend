package services

import (
	"context"
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/src/dtos"
)

type InvoiceService struct {
}

type InvoiceServiceDependencies struct {
	dig.In
}

func NewInvoiceService(deps InvoiceServiceDependencies) *InvoiceService {
	return &InvoiceService{}
}

func (i InvoiceService) CreateInvoice(
	ctx context.Context,
	input dtos.CreateInvoiceInput,
) (output dtos.CreateInvoiceOutput, err error) {
	return
}

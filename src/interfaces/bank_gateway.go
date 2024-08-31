package interfaces

import (
	"context"
	"joaosalless/challenge-starkbank/src/dtos"
)

type BankGateway interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
	CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error)
}

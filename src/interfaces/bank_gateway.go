package interfaces

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

type BankGateway interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
	CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error)
}

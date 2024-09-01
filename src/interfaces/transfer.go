package interfaces

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

type TransferService interface {
	CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error)
	CreateTransferFromInvoice(ctx context.Context, input dtos.CreateTransferFromInvoiceInput) (output dtos.CreateTransferOutput, err error)
}

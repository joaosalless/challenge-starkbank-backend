package interfaces

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
}

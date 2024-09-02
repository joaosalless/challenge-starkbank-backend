package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

type InvoiceHandler interface {
	CreateInvoice(ctx *gin.Context)
}

type InvoiceController interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
}

type InvoiceService interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
}

type InvoiceCreateCron interface {
	Run()
}

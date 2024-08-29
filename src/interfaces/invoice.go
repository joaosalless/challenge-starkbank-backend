package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"joaosalless/challenge-starkbank/src/dtos"
)

type InvoiceHandler interface {
	CreateInvoice(ctx *gin.Context)
	InvoiceHookProcess(ctx *gin.Context)
}

type InvoiceController interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
	InvoiceHookProcess(ctx context.Context, input dtos.InvoiceHookProcessInput) (output dtos.InvoiceHookProcessOutput, err error)
}

type InvoiceService interface {
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
}

type InvoiceCreateCron interface {
	Run()
}

package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceHandler struct {
	invoiceController interfaces.InvoiceController
}

type InvoiceHandlerDependencies struct {
	dig.In
	InvoiceController interfaces.InvoiceController `name:"InvoiceController"`
}

func NewInvoiceHandler(deps InvoiceHandlerDependencies) *InvoiceHandler {
	return &InvoiceHandler{invoiceController: deps.InvoiceController}
}

func (h InvoiceHandler) CreateInvoice(ctx *gin.Context) {
	var input dtos.CreateInvoiceInput

	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	invoice, err := h.invoiceController.CreateInvoice(ctx, input)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, invoice)
}

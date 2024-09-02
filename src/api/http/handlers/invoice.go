package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type InvoiceHandler struct {
	app               interfaces.Application
	invoiceController interfaces.InvoiceController
}

type InvoiceHandlerDependencies struct {
	ioc.In
	Application       interfaces.Application       `name:"Application"`
	InvoiceController interfaces.InvoiceController `name:"InvoiceController"`
}

func NewInvoiceHandler(deps InvoiceHandlerDependencies) *InvoiceHandler {
	return &InvoiceHandler{app: deps.Application, invoiceController: deps.InvoiceController}
}

type CreateInvoiceRequestBody struct {
	Invoices []domain.Invoice `json:"invoices" binding:"required"`
}

func (h InvoiceHandler) CreateInvoice(ctx *gin.Context) {
	var body CreateInvoiceRequestBody

	err := json.NewDecoder(ctx.Request.Body).Decode(&body)
	if err != nil {
		h.app.Logger().Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	invoices, err := h.invoiceController.CreateInvoice(ctx, dtos.CreateInvoiceInput{Data: body.Invoices})
	if err != nil {
		h.app.Logger().Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	h.app.Logger().Infow(fmt.Sprintf("Invoice created successfully: %+v", invoices))
	ctx.JSON(201, invoices)
}

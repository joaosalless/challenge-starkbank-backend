package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/app"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"time"
)

type InvoiceHandler struct {
	logger            interfaces.Logger
	invoiceController interfaces.InvoiceController
}

type InvoiceHandlerDependencies struct {
	app.Dependencies
	InvoiceController interfaces.InvoiceController `name:"InvoiceController"`
}

func NewInvoiceHandler(deps InvoiceHandlerDependencies) *InvoiceHandler {
	return &InvoiceHandler{logger: deps.Logger, invoiceController: deps.InvoiceController}
}

type CreateInvoiceRequestBody struct {
	Invoices []domain.Invoice `json:"invoices" binding:"required"`
}

func (h InvoiceHandler) CreateInvoice(ctx *gin.Context) {
	var body CreateInvoiceRequestBody

	err := json.NewDecoder(ctx.Request.Body).Decode(&body)
	if err != nil {
		h.logger.Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	invoices, err := h.invoiceController.CreateInvoice(ctx, dtos.CreateInvoiceInput{Data: body.Invoices})
	if err != nil {
		h.logger.Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infow(fmt.Sprintf("Invoice created successfully: %+v", invoices))
	ctx.JSON(201, invoices)
}

type InvoiceHookRequestBody struct {
	Event struct {
		Created time.Time `json:"created"`
		Id      string    `json:"id"`
		Log     struct {
			Created time.Time     `json:"created"`
			Errors  []interface{} `json:"errors"`
			Id      string        `json:"id"`
			Invoice struct {
				Amount       int       `json:"amount"`
				Brcode       string    `json:"brcode"`
				Created      time.Time `json:"created"`
				Descriptions []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"descriptions"`
				DiscountAmount int `json:"discountAmount"`
				Discounts      []struct {
					Due        time.Time `json:"due"`
					Percentage float64   `json:"percentage"`
				} `json:"discounts"`
				Due            time.Time     `json:"due"`
				Expiration     int           `json:"expiration"`
				Fee            int           `json:"fee"`
				Fine           float64       `json:"fine"`
				FineAmount     int           `json:"fineAmount"`
				Id             string        `json:"id"`
				Interest       float64       `json:"interest"`
				InterestAmount int           `json:"interestAmount"`
				Link           string        `json:"link"`
				Name           string        `json:"name"`
				NominalAmount  int           `json:"nominalAmount"`
				Pdf            string        `json:"pdf"`
				Rules          []interface{} `json:"rules"`
				Splits         []interface{} `json:"splits"`
				Status         string        `json:"status"`
				Tags           []string      `json:"tags"`
				TaxId          string        `json:"taxId"`
				TransactionIds []interface{} `json:"transactionIds"`
				Updated        time.Time     `json:"updated"`
			} `json:"invoice"`
			Type string `json:"type"`
		} `json:"log"`
		Subscription string `json:"subscription"`
		WorkspaceId  string `json:"workspaceId"`
	} `json:"event"`
}

func (body InvoiceHookRequestBody) ParseInvoiceEvent() (domain.InvoiceWebhookEvent, error) {
	if body.Event.Subscription != "invoice" {
		return domain.InvoiceWebhookEvent{}, errors.New("unsupported resource")
	}

	event := domain.InvoiceWebhookEvent{
		ID:      body.Event.Id,
		Type:    body.Event.Log.Type,
		Created: body.Event.Log.Created,
		Log: domain.InvoiceWebhookEventLog{
			ID:      body.Event.Log.Id,
			Created: time.Time{},
			Errors:  nil,
			Invoice: domain.Invoice{
				Id:             body.Event.Log.Invoice.Id,
				Amount:         body.Event.Log.Invoice.Amount,
				Name:           body.Event.Log.Invoice.Name,
				TaxId:          body.Event.Log.Invoice.TaxId,
				Due:            &body.Event.Log.Invoice.Due,
				Expiration:     body.Event.Log.Invoice.Expiration,
				Fine:           body.Event.Log.Invoice.Fine,
				Interest:       body.Event.Log.Invoice.Interest,
				Tags:           body.Event.Log.Invoice.Tags,
				Pdf:            body.Event.Log.Invoice.Pdf,
				Link:           body.Event.Log.Invoice.Link,
				NominalAmount:  body.Event.Log.Invoice.NominalAmount,
				FineAmount:     body.Event.Log.Invoice.FineAmount,
				InterestAmount: body.Event.Log.Invoice.InterestAmount,
				DiscountAmount: body.Event.Log.Invoice.DiscountAmount,
				Brcode:         body.Event.Log.Invoice.Brcode,
				Status:         body.Event.Log.Invoice.Status,
				Fee:            body.Event.Log.Invoice.Fee,
			},
		},
	}

	return event, nil
}

func (h InvoiceHandler) HookProcessInvoice(ctx *gin.Context) {
	var body InvoiceHookRequestBody

	err := json.NewDecoder(ctx.Request.Body).Decode(&body)
	if err != nil {
		h.logger.Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payload, err := body.ParseInvoiceEvent()
	if err != nil {
		h.logger.Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	invoice, err := h.invoiceController.InvoiceHookProcess(ctx, dtos.InvoiceHookProcessInput{
		Event: payload,
	})
	if err != nil {
		h.logger.Errorw(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, invoice)
}

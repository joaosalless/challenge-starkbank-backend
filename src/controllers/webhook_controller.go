package controllers

import (
	"context"
	"encoding/json"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	StarkbankEvent "github.com/starkbank/sdk-go/starkbank/event"
	StarkbankInvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
)

type WebhookController struct {
	app             interfaces.Application
	bankGateway     interfaces.BankGateway
	invoiceService  interfaces.InvoiceService
	transferService interfaces.TransferService
}

type WebhookControllerDependencies struct {
	ioc.In
	Application     interfaces.Application     `name:"Application"`
	BankGateway     interfaces.BankGateway     `name:"BankGateway"`
	InvoiceService  interfaces.InvoiceService  `name:"InvoiceService"`
	TransferService interfaces.TransferService `name:"TransferService"`
}

func NewWebhookController(deps WebhookControllerDependencies) *WebhookController {
	return &WebhookController{
		app:             deps.Application,
		bankGateway:     deps.BankGateway,
		invoiceService:  deps.InvoiceService,
		transferService: deps.TransferService,
	}
}

func (wc WebhookController) ProcessEvent(ctx context.Context, input dtos.WebhookProcessEventInput) (output dtos.WebhookProcessEventOutput, err error) {
	wc.app.Logger().Infow("WebhookController.ProcessEvents called", "input", input)

	parsedEventInterface := StarkbankEvent.Parse(string(input.Content), input.Signature, wc.bankGateway.GetUser())
	if parsedEventInterface == nil {
		wc.app.Logger().Errorw("Failed to parse webhook event")
		output.Errors = append(output.Errors, domain.Error{Code: "webhook.failed_to_parse_event", Message: "Failed to parse webhook event"})
		return output, err
	}

	parsedEventStr, ok := parsedEventInterface.(string)
	if !ok {
		wc.app.Logger().Errorw("Failed to parse webhook event")
		output.Errors = append(output.Errors, domain.Error{Code: "webhook.failed_to_parse_event", Message: "Unexpected type returned from StarkbankEvent.Parse"})
		return output, err
	}

	var eventContainer struct {
		Event StarkbankEvent.Event `json:"event"`
	}

	err = json.Unmarshal([]byte(parsedEventStr), &eventContainer)
	if err != nil {
		wc.app.Logger().Errorw("Failed to unmarshal event JSON", "error", err)
		output.Errors = append(output.Errors, domain.Error{Code: "webhook.failed_to_parse_event", Message: "Failed to unmarshal event JSON"})
		return output, err
	}

	parsedEvent := eventContainer.Event.ParseLog()

	parsedLog, ok := parsedEvent.Log.(StarkbankInvoiceLog.Log)
	if !ok {
		wc.app.Logger().Errorw("Failed to parse webhook event log")
		output.Errors = append(output.Errors, domain.Error{Code: "webhook.failed_to_parse_event_log", Message: "Failed to parse webhook event log"})
		return output, err
	}

	if parsedEvent.Subscription == "invoice" {
		parsedInvoice := parsedLog.Invoice

		if parsedInvoice.Status == "paid" {
			_, err = wc.transferService.CreateTransferFromInvoice(ctx, dtos.CreateTransferFromInvoiceInput{Data: parsedInvoice})
			if err != nil {
				wc.app.Logger().Errorw("Failed to CreateTransferFromInvoice", "error", err)
				return output, err
			}
			return output, nil
		}
	}

	output.Message = "Webhook event log processed successfully"

	return output, nil
}

package controllers

import (
	"context"
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
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

	event, err := wc.bankGateway.ParseEvent(ctx, input)
	if err != nil {
		return output, fmt.Errorf("failed to parse event: %w", err)
	}

	if event.Subscription == domain.EventSubscriptionInvoice {
		eventLog, err := wc.bankGateway.ParseInvoiceEventLog(ctx, event)
		if err != nil {
			return output, fmt.Errorf("failed to parse %s event log: %w", event.Subscription, err)
		}

		wc.app.Logger().Infow("transfer event log parsed successfully", "eventLog", eventLog)

		if eventLog.Type == domain.InvoiceEventCredited {
			_, err := wc.transferService.CreateTransferFromInvoice(ctx, dtos.CreateTransferFromInvoiceInput{Data: domain.Invoice(eventLog.Invoice)})
			if err != nil {
				return output, err
			}
		}
	} else if event.Subscription == domain.EventSubscriptionTransfer {
		eventLog, err := wc.bankGateway.ParseTransferEventLog(ctx, event)
		if err != nil {
			return output, fmt.Errorf("failed to parse %s event log: %w", event.Subscription, err)
		}

		wc.app.Logger().Infow("transfer event log parsed successfully", "eventLog", eventLog)
	}

	output.Message = "Webhook event log processed successfully"

	return output, nil
}

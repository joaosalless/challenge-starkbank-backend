package dtos

import (
	"joaosalless/challenge-starkbank/src/domain"
)

type CreateInvoiceInput struct {
	Data []domain.Invoice `json:",omitempty"`
}

type CreateInvoiceOutput struct {
	Data   []domain.Invoice `json:",omitempty"`
	Errors []domain.Error   `json:",omitempty"`
}

type InvoiceHookProcessInput struct {
	Event domain.InvoiceWebhookEvent `json:",omitempty"`
}

type InvoiceHookProcessOutput struct {
	Message string         `json:",omitempty"`
	Errors  []domain.Error `json:",omitempty"`
}

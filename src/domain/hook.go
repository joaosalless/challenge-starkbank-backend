package domain

import "time"

type ResourceType string

var (
	ResourceTypeBoleto   ResourceType = "Boleto"
	ResourceTypeInvoice  ResourceType = "Invoice"
	ResourceTypeTransfer ResourceType = "Transfer"
)

type Webhook struct {
	Event    WebhookEvent `json:"event,omitempty"`
	Resource ResourceType `json:"resource,omitempty"`
}

type WebhookEvent struct {
	Type    string    `json:"type,omitempty"`
	ID      string    `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type WebhookEventLog struct {
	ID      string    `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type InvoiceWebhook struct {
	Resource ResourceType        `json:"resource,omitempty"`
	Event    InvoiceWebhookEvent `json:"event,omitempty"`
}

type InvoiceWebhookEvent struct {
	Type    string                 `json:"type,omitempty"`
	ID      string                 `json:"id,omitempty"`
	Created time.Time              `json:"created,omitempty"`
	Log     InvoiceWebhookEventLog `json:"log,omitempty"`
}

type InvoiceWebhookEventLog struct {
	ID      string    `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Errors  []Error   `json:"errors,omitempty"`
	Invoice Invoice   `json:"invoice,omitempty"`
}

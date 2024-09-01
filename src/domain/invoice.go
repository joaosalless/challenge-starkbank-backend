package domain

import (
	"time"
)

const (
	InvoiceStatusCreated = "created"
	InvoiceStatusPaid    = "paid"

	InvoiceEventCreated  = "created"
	InvoiceEventCredited = "credited"
)

type Invoice struct {
	Id                 string                   `json:",omitempty"`
	Amount             int                      `json:",omitempty"`
	Name               string                   `json:",omitempty"`
	TaxId              string                   `json:",omitempty"`
	Due                *time.Time               `json:",omitempty"`
	Expiration         int                      `json:",omitempty"`
	Fine               float64                  `json:",omitempty"`
	Interest           float64                  `json:",omitempty"`
	Discounts          []map[string]interface{} `json:",omitempty"`
	Tags               []string                 `json:",omitempty"`
	Rules              []RuleStrings            `json:",omitempty"`
	Descriptions       []map[string]interface{} `json:",omitempty"`
	DisplayDescription string                   `json:",omitempty"`
	Pdf                string                   `json:",omitempty"`
	Link               string                   `json:",omitempty"`
	NominalAmount      int                      `json:",omitempty"`
	FineAmount         int                      `json:",omitempty"`
	InterestAmount     int                      `json:",omitempty"`
	DiscountAmount     int                      `json:",omitempty"`
	Brcode             string                   `json:",omitempty"`
	Status             string                   `json:",omitempty"`
	Fee                int                      `json:",omitempty"`
	TransactionIds     []string                 `json:",omitempty"`
	Created            *time.Time               `json:",omitempty"`
	Updated            *time.Time               `json:",omitempty"`
	ExternalId         string                   `json:",omitempty"`
}

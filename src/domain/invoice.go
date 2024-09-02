package domain

import (
	StarkbankInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
)

const (
	InvoiceStatusCreated = "created"
	InvoiceStatusPaid    = "paid"

	InvoiceEventCreated  = "created"
	InvoiceEventCredited = "credited"
)

type Invoice StarkbankInvoice.Invoice

package domain

import (
	StarkbankInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
)

const InvoiceStatusPaid = "paid"

type Invoice StarkbankInvoice.Invoice

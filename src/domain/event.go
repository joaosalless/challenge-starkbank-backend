package domain

import (
	StarkbankEvent "github.com/starkbank/sdk-go/starkbank/event"
	StarkbankInvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	StarkbankTransferLog "github.com/starkbank/sdk-go/starkbank/transfer/log"
)

const (
	EventSubscriptionInvoice  = "invoice"
	EventSubscriptionTransfer = "transfer"

	InvoiceEventCredited = "credited"
)

type Event StarkbankEvent.Event
type InvoiceEventLog StarkbankInvoiceLog.Log
type TransferEventLog StarkbankTransferLog.Log

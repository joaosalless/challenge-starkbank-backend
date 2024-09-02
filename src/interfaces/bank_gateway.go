package interfaces

import (
	"context"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/starkbank/ecdsa-go/v2/ellipticcurve/privatekey"
)

type BankGateway interface {
	GetUser() BankGatewayUser
	CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error)
	CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error)
	ParseEvent(ctx context.Context, input dtos.WebhookProcessEventInput) (output domain.Event, err error)
	ParseInvoiceEventLog(ctx context.Context, event domain.Event) (log domain.InvoiceEventLog, err error)
	ParseTransferEventLog(ctx context.Context, event domain.Event) (log domain.TransferEventLog, err error)
}

type BankGatewayUser interface {
	GetAcessId() string
	GetEnvironment() string
	GetPrivateKey() *privatekey.PrivateKey
}

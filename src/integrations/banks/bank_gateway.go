package banks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	StarkbankEvent "github.com/starkbank/sdk-go/starkbank/event"
	StarkbankInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
	StarkbankInvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	StarkTransfer "github.com/starkbank/sdk-go/starkbank/transfer"
	StarkbankTransferLog "github.com/starkbank/sdk-go/starkbank/transfer/log"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

type BankGateway struct {
	app  interfaces.Application
	user interfaces.BankGatewayUser
}

type BankGatewayDependencies struct {
	ioc.In
	Config      *config.Config         `name:"Config"`
	Application interfaces.Application `name:"Application"`
}

func NewBankGateway(deps BankGatewayDependencies) *BankGateway {
	return &BankGateway{
		app: deps.Application,
		user: project.Project{
			Id:          deps.Config.BankProvider.Starkbank.ProjectId,
			Environment: deps.Config.BankProvider.Starkbank.Environment,
			PrivateKey:  deps.Config.BankProvider.Starkbank.PrivateKey,
		},
	}
}

func (g BankGateway) ParseEvent(ctx context.Context, input dtos.WebhookProcessEventInput) (output domain.Event, err error) {
	eventInterface := StarkbankEvent.Parse(string(input.Content), input.Signature, g.user)
	if eventInterface == nil {
		return output, errors.New("failed to parse webhook event")
	}

	parsedEventStr, ok := eventInterface.(string)
	if !ok {
		return output, errors.New("failed to parse webhook event interface")
	}

	var eventContainer struct {
		Event StarkbankEvent.Event `json:"event"`
	}

	err = json.Unmarshal([]byte(parsedEventStr), &eventContainer)
	if err != nil {
		return output, errors.New("failed to unmarshal event json")
	}

	parsedEvent := eventContainer.Event.ParseLog()

	return domain.Event(parsedEvent), nil
}

func (g BankGateway) ParseInvoiceEventLog(ctx context.Context, event domain.Event) (log domain.InvoiceEventLog, err error) {
	parsedLog, ok := event.Log.(StarkbankInvoiceLog.Log)
	if !ok {
		return domain.InvoiceEventLog{}, errors.New("failed to parse webhook invoice event log")
	}

	return domain.InvoiceEventLog(parsedLog), nil
}

func (g BankGateway) ParseTransferEventLog(ctx context.Context, event domain.Event) (log domain.TransferEventLog, err error) {
	parsedLog, ok := event.Log.(StarkbankTransferLog.Log)
	if !ok {
		return domain.TransferEventLog{}, errors.New("failed to parse webhook transfer event log")
	}

	return domain.TransferEventLog(parsedLog), nil
}

func (g BankGateway) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	var invoices []StarkbankInvoice.Invoice
	for _, item := range input.Data {
		invoices = append(invoices, StarkbankInvoice.Invoice(item))
	}

	starkbankInvoices, stackErr := StarkbankInvoice.Create(invoices, g.user)
	if stackErr.Errors != nil {
		g.app.Logger().Errorw("Failed to create invoices", "input", input, "stackErr", stackErr)
		return output, fmt.Errorf("error creating invoices: %+v", stackErr)
	}

	for _, item := range starkbankInvoices {
		output.Data = append(output.Data, domain.Invoice(item))
	}

	return output, err
}

func (g BankGateway) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	g.app.Logger().Infow("WebhookController.ProcessEvents called", "input", input)

	var transfers []StarkTransfer.Transfer
	for _, t := range input.Data {
		transfers = append(transfers, StarkTransfer.Transfer(t))
	}

	starkbankTransfers, stackErr := StarkTransfer.Create(transfers, g.user)
	if stackErr.Errors != nil {
		g.app.Logger().Errorw("Failed to create transfers", "input", input, "stackErr", stackErr)
		return output, fmt.Errorf("error creating transfers: %+v", stackErr)
	}

	for _, t := range starkbankTransfers {
		output.Data = append(output.Data, domain.Transfer(t))
	}

	return output, err
}

package bank

import (
	"context"
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	StarkbankInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
	StarkTransfer "github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

// BankGateway implements interfaces.BankGateway
type BankGateway struct {
	app  interfaces.Application
	user interfaces.BankGatewayUser
}

type Dependencies struct {
	ioc.In
	Config      *config.Config         `name:"Config"`
	Application interfaces.Application `name:"Application"`
}

func NewBankGateway(deps Dependencies) *BankGateway {
	return &BankGateway{
		app: deps.Application,
		user: project.Project{
			Id:          deps.Config.BankProvider.Starkbank.ProjectId,
			Environment: deps.Config.BankProvider.Starkbank.Environment,
			PrivateKey:  deps.Config.BankProvider.Starkbank.PrivateKey,
		},
	}
}

func (g BankGateway) GetUser() interfaces.BankGatewayUser {
	return g.user
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

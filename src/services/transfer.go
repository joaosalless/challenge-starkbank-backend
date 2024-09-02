package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type TransferService struct {
	app            interfaces.Application
	bankGateway    interfaces.BankGateway
	transferConfig config.Transfer
}

type TransferServiceDependencies struct {
	ioc.In
	Config      *config.Config         `name:"Config"`
	Application interfaces.Application `name:"Application"`
	BankGateway interfaces.BankGateway `name:"BankGateway"`
}

func NewTransferService(deps TransferServiceDependencies) *TransferService {
	return &TransferService{
		app:            deps.Application,
		bankGateway:    deps.BankGateway,
		transferConfig: deps.Config.Transfer,
	}
}

func (i TransferService) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	i.app.Logger().Infow("TransferService.CreateTransfer", "input", input)

	output, err = i.bankGateway.CreateTransfer(ctx, input)
	if err != nil {
		i.app.Logger().Errorw("failed to call bankGateway.CreateTransfer", "input", input)
		return dtos.CreateTransferOutput{}, fmt.Errorf("error when call bankGateway.CreateTransfer: %w", err)
	}

	return output, nil
}

func (i TransferService) CreateTransferFromInvoice(ctx context.Context, input dtos.CreateTransferFromInvoiceInput) (output dtos.CreateTransferOutput, err error) {
	i.app.Logger().Infow("TransferService.CreateTransferFromInvoice", "input", input)

	if input.Data.Status != domain.InvoiceStatusPaid {
		i.app.Logger().Errorw("invalid invoice status", "input", input)
		return output, errors.New("invalid invoice status")
	}

	transfer := domain.Transfer{
		Amount:        i.calculateTransferAmount(input.Data),
		Name:          i.transferConfig.BankAccount.Name,
		TaxId:         i.transferConfig.BankAccount.TaxId,
		BankCode:      i.transferConfig.BankAccount.BankCode,
		BranchCode:    i.transferConfig.BankAccount.BranchCode,
		AccountNumber: i.transferConfig.BankAccount.AccountNumber,
		AccountType:   i.transferConfig.BankAccount.AccountType,
		ExternalId:    fmt.Sprintf("invoice:%s", input.Data.Id),
		Tags: []string{
			fmt.Sprintf("invoice:%s", input.Data.Id),
		},
		Description: fmt.Sprintf("Payment for invoice #%s - %s", input.Data.Id, input.Data.DisplayDescription),
	}

	return i.CreateTransfer(ctx, dtos.CreateTransferInput{Data: []domain.Transfer{transfer}})
}

func (i TransferService) calculateTransferAmount(invoice domain.Invoice) int {
	baseAmount := invoice.NominalAmount - invoice.DiscountAmount
	finalAmount := baseAmount - invoice.Fee

	return finalAmount
}

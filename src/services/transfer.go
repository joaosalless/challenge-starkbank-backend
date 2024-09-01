package services

import (
	"context"
	"errors"
	"fmt"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/domain"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type TransferService struct {
	logger         interfaces.Logger
	bankGateway    interfaces.BankGateway
	transferConfig config.Transfer
}

type TransferServiceDependencies struct {
	app.Dependencies
	BankGateway interfaces.BankGateway `name:"BankGateway"`
}

func NewTransferService(deps TransferServiceDependencies) *TransferService {
	return &TransferService{
		logger:         deps.Logger,
		bankGateway:    deps.BankGateway,
		transferConfig: deps.Config.Transfer,
	}
}

func (i TransferService) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	i.logger.Infow("TransferService.CreateTransfer", "input", input)

	output, err = i.bankGateway.CreateTransfer(ctx, input)
	if err != nil {
		return dtos.CreateTransferOutput{}, fmt.Errorf("error when call bankGateway.CreateTransfer: %w", err)
	}

	return output, nil
}

func (i TransferService) CreateTransferFromInvoice(ctx context.Context, input dtos.CreateTransferFromInvoiceInput) (output dtos.CreateTransferOutput, err error) {
	i.logger.Infow("TransferService.CreateTransferFromInvoice", "input", input)

	if input.Data.Status != domain.InvoiceStatusPaid {
		i.logger.Errorw("invalid invoice status", "input", input)
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
		Tags: []string{
			fmt.Sprintf("urn:invoice:%s", input.Data.ExternalId),
		},
		Description: fmt.Sprintf("Payment for invoice #%s - %s", input.Data.ExternalId, input.Data.DisplayDescription),
	}

	return i.CreateTransfer(ctx, dtos.CreateTransferInput{Data: []domain.Transfer{transfer}})
}

func (i TransferService) calculateTransferAmount(invoice domain.Invoice) int {
	baseAmount := invoice.NominalAmount - invoice.DiscountAmount
	finalAmount := baseAmount - invoice.Fee

	return finalAmount
}

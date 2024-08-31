package services

import (
	"context"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/dtos"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type TransferService struct {
	logger      interfaces.Logger
	bankGateway interfaces.BankGateway
}

type TransferServiceDependencies struct {
	app.Dependencies
	BankGateway interfaces.BankGateway `name:"BankGateway"`
}

func NewTransferService(deps TransferServiceDependencies) *TransferService {
	return &TransferService{logger: deps.Logger, bankGateway: deps.BankGateway}
}

func (i TransferService) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	return
}

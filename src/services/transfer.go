package services

import (
	"context"
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/src/dtos"
)

type TransferService struct {
}

type TransferServiceDependencies struct {
	dig.In
}

func NewTransferService(deps TransferServiceDependencies) *TransferService {
	return &TransferService{}
}

func (i TransferService) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	return
}

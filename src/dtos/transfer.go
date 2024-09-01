package dtos

import "github.com/joaosalless/challenge-starkbank-backend/src/domain"

type CreateTransferInput struct {
	Data []domain.Transfer `json:",omitempty"`
}

type CreateTransferOutput struct {
	Data   []domain.Transfer `json:",omitempty"`
	Errors []domain.Error    `json:",omitempty"`
}

type CreateTransferFromInvoiceInput struct {
	Data domain.Invoice `json:",omitempty"`
}

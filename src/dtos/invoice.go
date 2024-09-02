package dtos

import (
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
)

type CreateInvoiceInput struct {
	Data []domain.Invoice `json:",omitempty"`
}

type CreateInvoiceOutput struct {
	Data   []domain.Invoice `json:",omitempty"`
	Errors []domain.Error   `json:",omitempty"`
}

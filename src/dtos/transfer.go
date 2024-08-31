package dtos

import "joaosalless/challenge-starkbank/src/domain"

type CreateTransferInput struct {
	Data []domain.Transfer `json:",omitempty"`
}

type CreateTransferOutput struct {
	Data   []domain.Transfer `json:",omitempty"`
	Errors []domain.Error    `json:",omitempty"`
}

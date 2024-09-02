package dtos

import "github.com/joaosalless/challenge-starkbank-backend/src/domain"

type WebhookProcessEventInput struct {
	Content   []byte `json:",omitempty"`
	Signature string `json:",omitempty"`
}

type WebhookProcessEventOutput struct {
	Message string         `json:",omitempty"`
	Errors  []domain.Error `json:",omitempty"`
}

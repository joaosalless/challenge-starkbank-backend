package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
)

type WebhookHandler interface {
	ProcessEvent(ctx *gin.Context)
}

type WebhookController interface {
	ProcessEvent(ctx context.Context, input dtos.WebhookProcessEventInput) (output dtos.WebhookProcessEventOutput, err error)
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"net/http"
)

type WebhookHandler struct {
	app               interfaces.Application
	webhookController interfaces.WebhookController
}

type WebhookHandlerDependencies struct {
	ioc.In
	Application       interfaces.Application       `name:"Application"`
	WebhookController interfaces.WebhookController `name:"WebhookController"`
}

func NewWebhookHandler(deps WebhookHandlerDependencies) *WebhookHandler {
	return &WebhookHandler{
		app:               deps.Application,
		webhookController: deps.WebhookController,
	}
}

func (h WebhookHandler) ProcessEvent(c *gin.Context) {
	h.app.Logger().Infow("WebhookHandler.ProcessEvents called")

	bodyBytes, err := c.GetRawData()
	if err != nil {
		h.app.Logger().Errorw("Failed to read request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	signature := c.GetHeader(h.app.Config().BankProvider.Starkbank.DigitalSignatureHeader)
	if signature == "" {
		h.app.Logger().Errorw("Missing signature header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing signature header"})
		return
	}

	output, err := h.webhookController.ProcessEvent(c, dtos.WebhookProcessEventInput{
		Content:   bodyBytes,
		Signature: signature,
	})
	if err != nil {
		h.app.Logger().Errorw(err.Error())
		c.JSON(400, gin.H{"errors": output.Errors})
		return
	}

	c.JSON(200, gin.H{"message": "Event processed successfully"})
}

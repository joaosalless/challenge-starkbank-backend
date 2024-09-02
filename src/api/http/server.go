package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type Server struct {
	app            interfaces.Application
	invoiceHandler interfaces.InvoiceHandler
	webhookHandler interfaces.WebhookHandler
}

type ServerDependencies struct {
	ioc.In
	Application    interfaces.Application    `name:"Application"`
	InvoiceHandler interfaces.InvoiceHandler `name:"InvoiceHandler"`
	WebhookHandler interfaces.WebhookHandler `name:"WebhookHandler"`
}

func NewServer(deps ServerDependencies) *Server {
	api := &Server{
		invoiceHandler: deps.InvoiceHandler,
		webhookHandler: deps.WebhookHandler,
	}

	r := gin.Default()

	v1 := r.Group("v1")
	v1.POST("/invoices", api.invoiceHandler.CreateInvoice)
	v1.POST("/hooks", api.webhookHandler.ProcessEvent)

	if err := r.Run(":" + deps.Application.Config().Api.Port); err != nil {
		panic(err)
	}

	return api
}

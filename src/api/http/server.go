package http

import (
	"github.com/gin-gonic/gin"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type Server struct {
	invoiceHandler interfaces.InvoiceHandler
}

type ServerDependencies struct {
	app.Dependencies
	InvoiceHandler interfaces.InvoiceHandler `name:"InvoiceHandler"`
}

func NewServer(deps ServerDependencies) *Server {
	api := &Server{
		invoiceHandler: deps.InvoiceHandler,
	}

	r := gin.Default()

	v1 := r.Group("v1")
	v1.POST("/invoices", api.invoiceHandler.CreateInvoice)
	v1.POST("/hooks/invoices", api.invoiceHandler.HookProcessInvoice)

	if err := r.Run(":" + deps.Config.Api.Port); err != nil {
		panic(err)
	}

	return api
}

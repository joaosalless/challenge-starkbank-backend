package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type Server struct {
	invoiceHandler interfaces.InvoiceHandler
}

type ServerDependencies struct {
	dig.In
	Cfg            *config.Config            `name:"Config"`
	InvoiceHandler interfaces.InvoiceHandler `name:"InvoiceHandler"`
}

func NewServer(deps ServerDependencies) *Server {
	app := &Server{
		invoiceHandler: deps.InvoiceHandler,
	}

	r := gin.Default()

	v1 := r.Group("v1")
	v1.POST("/invoices", app.invoiceHandler.CreateInvoice)
	v1.POST("/hooks/invoices", app.invoiceHandler.InvoiceHookProcess)

	if err := r.Run(":" + deps.Cfg.AppPort); err != nil {
		panic(err)
	}

	return app
}

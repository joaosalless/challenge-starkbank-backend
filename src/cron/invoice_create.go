package cron

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type InvoiceCreateCron struct {
	interval          string
	logger            interfaces.Logger
	invoiceController interfaces.InvoiceController
}

type InvoiceCronDependencies struct {
	dig.In
	Cfg               *config.Config               `name:"Config"`
	logger            interfaces.Logger            `name:"Logger"`
	InvoiceController interfaces.InvoiceController `name:"InvoiceController"`
}

func NewInvoiceCreateCron(deps InvoiceCronDependencies) *InvoiceCreateCron {
	return &InvoiceCreateCron{
		logger:            deps.logger,
		interval:          deps.Cfg.InvoiceCreateCronInterval,
		invoiceController: deps.InvoiceController,
	}
}

func (ic *InvoiceCreateCron) Run() {
	ic.logger.Infow("starting invoice cron", "interval", ic.interval)
}

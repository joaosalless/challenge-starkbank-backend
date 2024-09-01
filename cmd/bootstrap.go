package cmd

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/clock"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/logging"
	"github.com/joaosalless/challenge-starkbank-backend/src/api/http/handlers"
	"github.com/joaosalless/challenge-starkbank-backend/src/controllers"
	"github.com/joaosalless/challenge-starkbank-backend/src/gateways/bank"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"github.com/joaosalless/challenge-starkbank-backend/src/schedule"
	"github.com/joaosalless/challenge-starkbank-backend/src/services"
	"go.uber.org/dig"
)

func Initialize() *dig.Container {
	return ioc.New([]ioc.Dependency{
		{
			Constructor: config.LoadConfig,
			Interface:   nil,
			Name:        "Config",
		},
		{
			Constructor: clock.NewClock,
			Interface:   new(interfaces.Clock),
			Name:        "Clock",
		},
		{
			Constructor: logging.NewLogger,
			Interface:   new(interfaces.Logger),
			Name:        "Logger",
		},
		{
			Constructor: bank.NewBankGateway,
			Interface:   new(interfaces.BankGateway),
			Name:        "BankGateway",
		},
		{
			Constructor: services.NewInvoiceService,
			Interface:   new(interfaces.InvoiceService),
			Name:        "InvoiceService",
		},
		{
			Constructor: services.NewTransferService,
			Interface:   new(interfaces.TransferService),
			Name:        "TransferService",
		},
		{
			Constructor: controllers.NewInvoiceController,
			Interface:   new(interfaces.InvoiceController),
			Name:        "InvoiceController",
		},
		{
			Constructor: handlers.NewInvoiceHandler,
			Interface:   new(interfaces.InvoiceHandler),
			Name:        "InvoiceHandler",
		},
		{
			Constructor: schedule.NewInvoiceCreateScheduledTask,
			Interface:   new(interfaces.ScheduledTask),
			Name:        "InvoiceCreateScheduledTask",
		},
		{
			Constructor: schedule.NewScheduledTasks,
			Interface:   new(interfaces.ScheduledTasks),
			Name:        "ScheduledTasks",
		},
	})
}

package cmd

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/application"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/clock"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/logging"
	"github.com/joaosalless/challenge-starkbank-backend/src/api/http/handlers"
	"github.com/joaosalless/challenge-starkbank-backend/src/controllers"
	"github.com/joaosalless/challenge-starkbank-backend/src/integrations/banks"
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
			Constructor: application.New,
			Interface:   new(interfaces.Application),
			Name:        "Application",
		},
		{
			Constructor: banks.NewBankGateway,
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
			Constructor: handlers.NewWebhookHandler,
			Interface:   new(interfaces.WebhookHandler),
			Name:        "WebhookHandler",
		},
		{
			Constructor: controllers.NewWebhookController,
			Interface:   new(interfaces.WebhookController),
			Name:        "WebhookController",
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

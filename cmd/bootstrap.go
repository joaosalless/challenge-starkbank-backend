package cmd

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/pkg/clock"
	"joaosalless/challenge-starkbank/pkg/logging"
	"joaosalless/challenge-starkbank/src/api/http/handlers"
	"joaosalless/challenge-starkbank/src/controllers"
	"joaosalless/challenge-starkbank/src/gateways/bank"
	"joaosalless/challenge-starkbank/src/interfaces"
	"joaosalless/challenge-starkbank/src/schedule"
	"joaosalless/challenge-starkbank/src/services"
	"log"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Name        string
}

var container *dig.Container

func Initialize() *dig.Container {
	container = dig.New()

	deps := []Dependency{
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
	}

	for _, dep := range deps {
		var err error

		if dep.Interface != nil {
			err = container.Provide(dep.Constructor, dig.As(dep.Interface), dig.Name(dep.Name))
		} else {
			err = container.Provide(dep.Constructor, dig.Name(dep.Name))
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return container
}

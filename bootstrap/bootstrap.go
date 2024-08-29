package bootstrap

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/pkg/logging"
	"joaosalless/challenge-starkbank/src/controllers"
	"joaosalless/challenge-starkbank/src/http/handlers"
	"joaosalless/challenge-starkbank/src/interfaces"
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

	err := container.Provide(config.LoadConfig, dig.Name("Config"))
	if err != nil {
		log.Fatal(err)
	}

	deps := []Dependency{
		{
			Constructor: logging.NewLogger,
			Interface:   new(interfaces.Logger),
			Name:        "Logger",
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
	}

	for _, dep := range deps {
		err = container.Provide(dep.Constructor, dig.As(dep.Interface), dig.Name(dep.Name))
		if err != nil {
			log.Fatal(err)
		}
	}

	return container
}

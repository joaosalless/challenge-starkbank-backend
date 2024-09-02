package schedule

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"github.com/mvrilo/go-cpf"
	"time"
)

type InvoiceCreateScheduledTask struct {
	scheduledTime     string
	app               interfaces.Application
	invoiceController interfaces.InvoiceController
}

type InvoiceCreateScheduledTaskDependencies struct {
	ioc.In
	Config            *config.Config               `name:"Config"`
	Application       interfaces.Application       `name:"Application"`
	InvoiceController interfaces.InvoiceController `name:"InvoiceController"`
}

func NewInvoiceCreateScheduledTask(deps InvoiceCreateScheduledTaskDependencies) *InvoiceCreateScheduledTask {
	return &InvoiceCreateScheduledTask{
		app:               deps.Application,
		scheduledTime:     deps.Config.Scheduler.InvoiceCreateScheduledTime,
		invoiceController: deps.InvoiceController,
	}
}

func (ic *InvoiceCreateScheduledTask) Schedule() error {
	ic.app.Logger().Infow("initializing invoice schedule", "scheduledTime", ic.scheduledTime)

	return nil
}

func (ic *InvoiceCreateScheduledTask) Run() (err error) {
	ic.app.Logger().Infow("starting InvoiceCreateScheduledTask", "scheduledTime", ic.scheduledTime)

	var invoices []domain.Invoice

	for i := 0; i < gofakeit.Number(8, 12); i++ {
		due := gofakeit.DateRange(
			time.Now().AddDate(0, 0, 1),
			time.Now().AddDate(0, 0, 60))

		invoices = append(invoices, domain.Invoice{
			Amount:     gofakeit.Number(10000, 20000),
			Due:        &due,
			Expiration: 1,
			Name:       fmt.Sprintf("%s %s", gofakeit.Person().FirstName, gofakeit.Person().LastName),
			TaxId:      cpf.GeneratePretty(),
		})
	}

	ic.app.Logger().Infow("finished InvoiceCreateScheduledTask", "invoices", invoices)

	_, err = ic.invoiceController.CreateInvoice(context.Background(), dtos.CreateInvoiceInput{Data: invoices})
	if err != nil {
		ic.app.Logger().Errorw("error creating invoices", "error", err)
		return err
	}

	return nil
}

func (ic *InvoiceCreateScheduledTask) ScheduleName() string {
	return "ScheduledInvoiceCreate"
}

func (ic *InvoiceCreateScheduledTask) ScheduleTime() string {
	return ic.scheduledTime
}

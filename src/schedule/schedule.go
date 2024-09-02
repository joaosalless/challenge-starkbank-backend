package schedule

import (
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"github.com/robfig/cron/v3"
	"log"
)

type ScheduledTasks struct {
	app   interfaces.Application
	tasks []interfaces.ScheduledTask
}

type ScheduledTasksDependencies struct {
	ioc.In
	Config                     *config.Config           `name:"Config"`
	Application                interfaces.Application   `name:"Application"`
	InvoiceCreateScheduledTask interfaces.ScheduledTask `name:"InvoiceCreateScheduledTask"`
}

func NewScheduledTasks(deps ScheduledTasksDependencies) *ScheduledTasks {
	if !deps.Config.Scheduler.Enabled {
		log.Println("Scheduled tasks are disabled")
		return &ScheduledTasks{}
	}

	st := &ScheduledTasks{
		tasks: []interfaces.ScheduledTask{
			deps.InvoiceCreateScheduledTask,
		},
	}

	err := st.Run()

	if err != nil {
		log.Fatal("Failed to run scheduled tasks:", err)
	}

	return st
}

func (s ScheduledTasks) Run() error {
	for _, task := range s.tasks {
		s.schedule(task)
	}
	return nil
}

func (s ScheduledTasks) schedule(task interfaces.ScheduledTask) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(task.ScheduleTime(), func() {
		err := task.Run()
		if err != nil {
			return
		}
	})

	if err != nil {
		s.app.Logger().Errorw(fmt.Sprintf("Failed to schedule task: %s", task.ScheduleName()), err)
		panic(err)
	}

	c.Start()
	select {}
}

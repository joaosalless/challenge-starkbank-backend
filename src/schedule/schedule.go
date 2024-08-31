package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"joaosalless/challenge-starkbank/pkg/app"
	"joaosalless/challenge-starkbank/src/interfaces"
	"log"
)

type ScheduledTasks struct {
	logger interfaces.Logger
	tasks  []interfaces.ScheduledTask
}

type ScheduledTasksDependencies struct {
	app.Dependencies
	InvoiceCreateScheduledTask interfaces.ScheduledTask `name:"InvoiceCreateScheduledTask"`
}

func NewScheduledTasks(deps ScheduledTasksDependencies) *ScheduledTasks {
	if !deps.Cfg.Scheduler.Enabled {
		deps.Logger.Infow("Scheduler is disabled")
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
		s.logger.Errorw(fmt.Sprintf("Failed to schedule task: %s", task.ScheduleName()), err)
		panic(err)
	}

	c.Start()
	select {}
}

package interfaces

type ScheduledTask interface {
	Run() error
	ScheduleName() string
	ScheduleTime() string
}

type ScheduledTasks interface {
	Run() error
}

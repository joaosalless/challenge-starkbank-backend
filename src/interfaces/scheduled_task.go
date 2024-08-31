package interfaces

type ScheduledTask interface {
	Run() error
	Schedule() (err error)
	ScheduleName() string
	ScheduleTime() string
}

type ScheduledTasks interface {
	Run() error
}

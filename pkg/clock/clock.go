package clock

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"time"
)

type Clock struct {
	location *time.Location
}

type Dependencies struct {
	dig.In
	Cfg *config.Config `name:"Config"`
}

var location *time.Location

func NewClock(deps Dependencies) *Clock {
	location, _ = time.LoadLocation(deps.Cfg.Clock.Location)

	return &Clock{
		location: location,
	}
}

func (c Clock) Now() time.Time {
	return time.Now().In(c.location)
}

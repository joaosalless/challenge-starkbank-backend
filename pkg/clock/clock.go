package clock

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"time"
)

type Clock struct {
	location *time.Location
}

type Dependencies struct {
	ioc.In
	Config *config.Config `name:"Config"`
}

var location *time.Location

func NewClock(deps Dependencies) *Clock {
	location, _ = time.LoadLocation(deps.Config.Clock.Location)

	return &Clock{
		location: location,
	}
}

func (c Clock) Now() time.Time {
	return time.Now().In(c.location)
}

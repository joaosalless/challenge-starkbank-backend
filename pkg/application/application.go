package application

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type Application struct {
	clock  interfaces.Clock
	config *config.Config
	logger interfaces.Logger
}

type Dependencies struct {
	ioc.In
	Clock  interfaces.Clock  `name:"Clock"`
	Config *config.Config    `name:"Config"`
	Logger interfaces.Logger `name:"Logger"`
}

func New(deps Dependencies) *Application {
	return &Application{
		clock:  deps.Clock,
		config: deps.Config,
		logger: deps.Logger,
	}
}

func (a Application) Clock() interfaces.Clock {
	return a.clock
}

func (a Application) Config() *config.Config {
	return a.config
}

func (a Application) Logger() interfaces.Logger {
	return a.logger
}

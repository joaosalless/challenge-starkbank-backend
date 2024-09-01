package app

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
	"go.uber.org/dig"
)

type Dependencies struct {
	dig.In
	Config *config.Config    `name:"Config"`
	Clock  interfaces.Clock  `name:"Clock"`
	Logger interfaces.Logger `name:"Logger"`
}

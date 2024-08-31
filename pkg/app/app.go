package app

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type Dependencies struct {
	dig.In
	Config *config.Config    `name:"Config"`
	Clock  interfaces.Clock  `name:"Clock"`
	Logger interfaces.Logger `name:"Logger"`
}

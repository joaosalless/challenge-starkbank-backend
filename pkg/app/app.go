package app

import (
	"go.uber.org/dig"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type Dependencies struct {
	dig.In
	Cfg    *config.Config    `name:"Config"`
	Clock  interfaces.Clock  `name:"Clock"`
	Logger interfaces.Logger `name:"Logger"`
}

package app

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/ioc"
	"github.com/joaosalless/challenge-starkbank-backend/src/interfaces"
)

type Dependencies struct {
	ioc.In
	Config *config.Config    `name:"Config"`
	Clock  interfaces.Clock  `name:"Clock"`
	Logger interfaces.Logger `name:"Logger"`
}

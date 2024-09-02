package interfaces

import (
	"github.com/joaosalless/challenge-starkbank-backend/config"
)

type Application interface {
	Clock() Clock
	Config() *config.Config
	Logger() Logger
}

package logging

import (
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"joaosalless/challenge-starkbank/config"
	"joaosalless/challenge-starkbank/src/interfaces"
)

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

type Dependencies struct {
	dig.In
	Config *config.Config   `name:"Config"`
	Clock  interfaces.Clock `name:"Clock"`
}

func NewLogger(deps Dependencies) *Logger {
	zapConfig := zap.NewProductionConfig()

	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		sugaredLogger: logger.Sugar(),
	}
}

func (z *Logger) Infow(msg string, keysAndValues ...interface{}) {
	z.sugaredLogger.Infow(msg, keysAndValues...)
}

func (z *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	z.sugaredLogger.Warnw(msg, keysAndValues...)
}

func (z *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	z.sugaredLogger.Errorw(msg, keysAndValues...)
}

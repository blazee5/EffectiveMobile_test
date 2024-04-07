package logger

import (
	"go.uber.org/zap"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func NewLogger() *zap.SugaredLogger {
	var log *zap.Logger

	switch os.Getenv("ENV") {
	case envLocal:
		log, _ = zap.NewDevelopment()
	case envDev:
		log, _ = zap.NewDevelopment()
	case envProd:
		log, _ = zap.NewProduction()
	}

	defer log.Sync()

	return log.Sugar()
}

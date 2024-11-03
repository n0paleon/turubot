package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"turubot/infra/config"
	"turubot/internal/domain"
)

type ConsoleLogger struct {
	*zap.SugaredLogger
}

var C *zap.SugaredLogger

func NewConsoleLogger(cfg *config.Config) domain.Logger {
	logPath := []string{"stdout"}
	errPath := []string{"stderr"}

	logOpts := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(cfg.Log.Level)),
		Encoding:         cfg.Log.Encoding,
		Development:      cfg.App.Development,
		OutputPaths:      logPath,
		ErrorOutputPaths: errPath,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	logger, err := logOpts.Build()
	if err != nil {
		return nil
	}

	sugaredLogger := logger.Sugar()
	C = sugaredLogger

	return C
}

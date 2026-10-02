package log

import (
	"github.com/AliAstanov/my_blogs/config"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Log struct {
	logger.LoggerI
}

func (l *Log) Cleanup() {
	if err := logger.Cleanup(l); err != nil {
		l.Error("Failed to cleanup logger", logger.Error(err))
	}
}

func NewLogger(cfg config.GeneralConfig) Log {
	var loggerLewel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLewel = logger.LevelDebug
	case config.TestMode:
		loggerLewel = logger.LevelDebug
	default:
		loggerLewel = logger.LevelInfo
	}

	logger := logger.NewLogger(cfg.AppName, loggerLewel)

	return Log{logger}
}

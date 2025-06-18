package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init(level string) error {
	var err error
	once.Do(func() {
		parsedLevel := zapcore.InfoLevel
		if err := parsedLevel.UnmarshalText([]byte(level)); err != nil {
			parsedLevel = zapcore.InfoLevel
		}
		cfg := zap.Config{
			Level:            zap.NewAtomicLevelAt(parsedLevel),
			Encoding:         "console",
			OutputPaths:      []string{"logs/debug.log"},
			ErrorOutputPaths: []string{"logs/debug.log"},
			EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		}

		logger, err = cfg.Build()

		logger.Debug("Logger initialzed.")
	})

	return err
}

func L() *zap.Logger {
	if logger == nil {
		panic("call logger init first!!!")
	}

	return logger
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

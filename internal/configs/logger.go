package configs

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(lc fx.Lifecycle, envs Environments) *zap.Logger {
	var logger *zap.Logger

	switch envs.MODE.String() {
	case "production":
		logger, _ = zap.NewProduction()
	default:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.OutputPaths = []string{"stderr"}
		logger, _ = config.Build()
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger
}

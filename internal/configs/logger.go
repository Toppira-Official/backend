package configs

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle, envs Environments) *zap.Logger {
	var logger *zap.Logger

	switch envs.MODE.String() {
	case "production":
		logger, _ = zap.NewProduction()
	default:
		logger, _ = zap.NewDevelopment()
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger
}

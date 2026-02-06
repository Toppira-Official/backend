package main

import (
	"github.com/Toppira-Official/backend/internal/configs"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.
		New(
			fx.Provide(
				configs.GetEnvironments,
				configs.NewHttpServer,
				configs.NewLogger,
			),
			fx.Invoke(
				configs.LoadEnvironmentsFromEnvFile,
			),
			fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			}),
		).
		Run()
}

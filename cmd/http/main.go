package main

import (
	"github.com/Toppira-Official/backend/internal/configs"
	"github.com/Toppira-Official/backend/internal/scripts"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func init() {
	configs.LoadEnvironmentsFromEnvFile()
}

func main() {
	fx.
		New(
			fx.Provide(
				configs.GetEnvironments,
				configs.NewHttpServer,
				configs.NewLogger,
				configs.NewDB,
				configs.NewQuery,
			),
			fx.Invoke(
				scripts.LoadMigrations,
			),
			fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			}),
		).
		Run()
}

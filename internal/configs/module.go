package configs

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"configs",
	fx.Provide(
		GetEnvironments,
		NewHttpServer,
		NewLogger,
		NewDB,
		NewQuery,
	),
	fx.Invoke(LoadMigrations),
)

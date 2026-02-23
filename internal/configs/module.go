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
		NewCache,
		NewQuery,
		GetGoogleOauthConfig,
		NewRateLimiter,
	),
	fx.Invoke(LoadMigrations),
)

package main

import (
	_ "github.com/Toppira-Official/backend/docs"
	"github.com/Toppira-Official/backend/internal/configs"
	"github.com/Toppira-Official/backend/internal/modules/auth"
	"github.com/Toppira-Official/backend/internal/modules/user"
	"github.com/Toppira-Official/backend/internal/scripts"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func init() {
	configs.LoadEnvironmentsFromEnvFile()
}

//	@title			Toppira APIs
//	@version		1.0
//	@description	Toppira APIs Documents.

//	@contact.name	Ali Moradi
//	@contact.email	AliMoradi0Business@gmail.com

// @host	localhost:3000
func main() {
	fx.
		New(
			fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			}),
			fx.Invoke(func(r *gin.Engine) {
				r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			}),
			configs.Module,
			scripts.Module,
			user.Module,
			auth.Module,
		).
		Run()
}

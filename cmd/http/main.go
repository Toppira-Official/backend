package main

import (
	_ "github.com/Toppira-Official/Reminder_Server/docs"
	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/middlewares"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/queues"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/utils"
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

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

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
			user.Module,
			auth.Module,
			middlewares.Module,
			utils.Module,
			queues.Module,
		).
		Run()
}

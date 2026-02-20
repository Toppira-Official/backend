package user

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	GetMeHandler    *handler.GetMeHandler
	UpdateMeHandler *handler.UpdateMeHandler
	GuardLogin      gin.HandlerFunc `name:"guard_login"`
}

func RegisterRoutes(engine *gin.Engine, h Handler) {
	group := engine.Group("/user")

	group.GET("/me", h.GuardLogin, h.GetMeHandler.GetMyInfo)
	group.PATCH("/me", h.GuardLogin, h.UpdateMeHandler.UpdateMyInfo)
}

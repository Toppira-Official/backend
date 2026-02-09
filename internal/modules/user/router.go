package user

import (
	"github.com/Toppira-Official/backend/internal/modules/user/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UserHandlers struct {
	fx.In

	Me         *handler.GetMeHandler
	GuardLogin gin.HandlerFunc `name:"guard_login"`
}

func RegisterRoutes(engine *gin.Engine, h UserHandlers) {
	group := engine.Group("/user")

	group.GET("/me", h.GuardLogin, h.Me.GetMyInfo)
}

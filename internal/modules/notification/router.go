package notification

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	GuardLogin               gin.HandlerFunc `name:"guard_login"`
	SubscribeFirebaseHandler *handler.SubscribeFirebaseHandler
}

func RegisterRoutes(engine *gin.Engine, h Handler) {
	group := engine.Group("/notification")
	group.POST("/firebase/subscribe", h.GuardLogin, h.SubscribeFirebaseHandler.Subscribe)
}

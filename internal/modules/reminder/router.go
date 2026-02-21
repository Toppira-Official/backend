package reminder

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	GuardLogin         gin.HandlerFunc `name:"guard_login"`
	NewReminderHandler *handler.NewReminderHandler
}

func RegisterRoutes(engine *gin.Engine, h Handler) {
	group := engine.Group("/reminder")
	group.POST("/", h.GuardLogin, h.NewReminderHandler.NewReminder)
}

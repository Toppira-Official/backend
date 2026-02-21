package reminder

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"reminder",
	fx.Provide(
		usecase.NewCreateReminderUsecase,
	),
)

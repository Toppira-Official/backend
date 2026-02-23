package reminder

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/validator"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"reminder",
	fx.Provide(
		usecase.NewCreateReminderUsecase,
		usecase.NewListRemindersUsecase,
		usecase.NewDeleteeReminderUsecase,
		handler.NewNewReminderHandler,
		handler.NewMyRemindersHandler,
	),
	fx.Invoke(
		RegisterRoutes,
		validator.RegisterPriorityValidators,
	),
)

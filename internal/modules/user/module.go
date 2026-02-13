package user

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/handler"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		handler.NewGetMeHandler,
		handler.NewUpdateMeHandler,
		usecase.NewCreateUserUsecase,
		usecase.NewUpdateUserUsecase,
		usecase.NewFindUserByEmailUsecase,
		usecase.NewFindUserByIDUsecase,
	),
	fx.Invoke(RegisterRoutes),
)

package user

import (
	"github.com/Toppira-Official/backend/internal/modules/user/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		usecase.NewCreateUserUsecase,
	),
)

package auth

import (
	"github.com/Toppira-Official/backend/internal/modules/auth/handler"
	"github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth",
	fx.Provide(
		handler.NewSignUpHandler,
		usecase.NewCreateUserUsecase,
		usecase.NewVerifyPasswordUsecase,
		usecase.NewGenerateJwtUsecase,
		usecase.NewVerifyJwtUsecase,
	),
	fx.Invoke(handler.RegisterRoutes),
)

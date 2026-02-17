package auth

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth",
	fx.Provide(
		handler.NewSignUpHandler,
		handler.NewLoginHandler,
		handler.NewGoogleOauthHandler,
		usecase.NewCreateUserUsecase,
		usecase.NewVerifyPasswordUsecase,
		usecase.NewGenerateJwtUsecase,
		usecase.NewVerifyJwtUsecase,
		usecase.NewGoogleOauthRedirectURLUsecase,
		usecase.NewGoogleOauthCallbackUsecase,
	),
	fx.Invoke(RegisterRoutes),
)

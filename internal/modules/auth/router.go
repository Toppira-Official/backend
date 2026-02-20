package auth

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	SignUp      *handler.SignUpHandler
	Login       *handler.LoginHandler
	GoogleOauth *handler.GoogleOauthHandler
}

func RegisterRoutes(engine *gin.Engine, h Handler) {
	group := engine.Group("/auth")

	group.GET("/google-oauth/redirect-url", h.GoogleOauth.GetGoogleOauthRedirectURL)
	group.GET("/google-oauth/callback", h.GoogleOauth.GoogleOauthCallback)
	group.POST("/sign-up-with-user-password", h.SignUp.SignUpWithEmailPassword)
	group.POST("/login-with-user-password", h.Login.LoginWithEmailPassword)
}

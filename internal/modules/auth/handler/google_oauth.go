package handler

import (
	"net/http"

	"github.com/Toppira-Official/backend/internal/configs"
	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	output "github.com/Toppira-Official/backend/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

const (
	googleOauthStateCookieName = "google_oauth_state"
)

type GoogleOauthHandler struct {
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase
	envs                          configs.Environments
}

func NewGoogleOauthHandler(
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase,
	envs configs.Environments,
) *GoogleOauthHandler {
	return &GoogleOauthHandler{
		googleOauthRedirectURLUsecase: googleOauthRedirectURLUsecase,
		envs:                          envs,
	}
}

// GetGoogleOauthRedirectURL godoc
//
//	@Summary	get google oauth redirect url
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	output.HttpOutput
//	@Router		/auth/google-oauth/redirect-url [get]
func (h *GoogleOauthHandler) GetGoogleOauthRedirectURL(c *gin.Context) {
	ctx := c.Request.Context()
	redirectUrl, state := h.googleOauthRedirectURLUsecase.Execute(ctx)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     googleOauthStateCookieName,
		Value:    state,
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   h.envs.MODE.String() == "production",
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, output.HttpOutput{
		Data: map[string]any{
			"redirect_url": redirectUrl,
		}},
	)
}

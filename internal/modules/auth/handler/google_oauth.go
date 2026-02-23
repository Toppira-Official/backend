package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler/dto"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	_ "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type GoogleOauthHandler struct {
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase
	googleOauthCallbackUsecase    authUsecase.GoogleOauthCallbackUsecase
	findUserByEmailUsecase        userUsecase.FindUserByEmailUsecase
	generateJwtUsecase            authUsecase.GenerateJwtUsecase
	createUserUsecase             userUsecase.CreateUserUsecase
	envs                          configs.Environments
}

func NewGoogleOauthHandler(
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase,
	googleOauthCallbackUsecase authUsecase.GoogleOauthCallbackUsecase,
	findUserByEmailUsecase userUsecase.FindUserByEmailUsecase,
	generateJwtUsecase authUsecase.GenerateJwtUsecase,
	createUserUsecase userUsecase.CreateUserUsecase,
	envs configs.Environments,
) *GoogleOauthHandler {
	return &GoogleOauthHandler{
		googleOauthRedirectURLUsecase: googleOauthRedirectURLUsecase,
		googleOauthCallbackUsecase:    googleOauthCallbackUsecase,
		findUserByEmailUsecase:        findUserByEmailUsecase,
		generateJwtUsecase:            generateJwtUsecase,
		createUserUsecase:             createUserUsecase,
		envs:                          envs,
	}
}

// GetGoogleOauthRedirectURL godoc
//
//	@Summary	Redirect to Google OAuth URL
//	@Tags		Authentication
//	@Success	307	{string}	string	Redirect	to	Google	OAuth	URL
//	@Failure	500	{object}	errors.ClientError
//	@Router		/auth/google-oauth/redirect-url [get]
func (h *GoogleOauthHandler) GetGoogleOauthRedirectURL(c *gin.Context) {
	ctx := c.Request.Context()
	redirectUrl, err := h.googleOauthRedirectURLUsecase.Execute(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

// GoogleOauthCallback godoc
//
//	@Summary	Handle Google OAuth callback
//	@Tags		Authentication
//	@Param		code	query		string	true	"Code"
//	@Param		state	query		string	true	"State"
//	@Success	200		{object}	output.HttpOutput
//	@Failure	401		{object}	errors.ClientError
//	@Failure	500		{object}	errors.ClientError
//	@Router		/auth/google-oauth/callback [get]
func (h *GoogleOauthHandler) GoogleOauthCallback(c *gin.Context) {
	ctx := c.Request.Context()

	code := c.Query("code")
	state := c.Query("state")

	userInfo, err := h.googleOauthCallbackUsecase.Execute(ctx, code, state)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := h.findUserByEmailUsecase.Execute(ctx, userInfo.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userName := strings.TrimSpace(fmt.Sprintf("%s %s", userInfo.Name, userInfo.FamilyName))
		user, err = h.createUserUsecase.Execute(ctx, &input.CreateUserInput{
			Email:          userInfo.Email,
			ProfilePicture: &userInfo.Picture,
			Name:           &userName,
		})
		if err != nil {
			c.Error(err)
			return
		}
	}

	userIDString := strconv.Itoa(int(user.ID))
	accessToken, err := h.generateJwtUsecase.Execute(ctx, userIDString)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output.HttpOutput[dto.GoogleOAuthOutput]{
		Data: dto.GoogleOAuthOutput{
			User:        userInfo,
			AccessToken: accessToken,
		},
	})
}

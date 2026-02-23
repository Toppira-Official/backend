package handler

import (
	"net/http"
	"strconv"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler/dto"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"

	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	verifyPasswordUsecase  authUsecase.VerifyPasswordUsecase
	findUserByEmailUsecase userUsecase.FindUserByEmailUsecase
	generateJwtUsecase     authUsecase.GenerateJwtUsecase
}

func NewLoginHandler(
	verifyPasswordUsecase authUsecase.VerifyPasswordUsecase,
	generateJwtUsecase authUsecase.GenerateJwtUsecase,
	findUserByEmailUsecase userUsecase.FindUserByEmailUsecase,
) *LoginHandler {
	return &LoginHandler{
		verifyPasswordUsecase:  verifyPasswordUsecase,
		generateJwtUsecase:     generateJwtUsecase,
		findUserByEmailUsecase: findUserByEmailUsecase,
	}
}

// LoginWithEmailPassword godoc
//
//	@Summary	login with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.LoginWithEmailPasswordInput	true	"Login Input"
//	@Success	200		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	404		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Router		/auth/login-with-user-password [post]
func (hl *LoginHandler) LoginWithEmailPassword(c *gin.Context) {
	ctx := c.Request.Context()

	var input dto.LoginWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	user, err := hl.findUserByEmailUsecase.Execute(ctx, input.Email)
	if err != nil {
		c.Error(err)
		return
	}

	isPasswordValid := hl.verifyPasswordUsecase.Execute(ctx, []byte(input.Password), []byte(*user.Password))
	if !isPasswordValid {
		c.Error(apperrors.E(apperrors.ErrAuthInvalidEmailOrPassword, err))
		return
	}

	userIDString := strconv.Itoa(int(user.ID))
	accessToken, err := hl.generateJwtUsecase.Execute(ctx, userIDString)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output.HttpOutput[dto.AuthOutput]{
		Data: dto.AuthOutput{
			User:        output.ToUserOutput(user),
			AccessToken: accessToken,
		},
	})
}

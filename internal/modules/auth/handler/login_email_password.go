package handler

import (
	"net/http"
	"strconv"

	input "github.com/Toppira-Official/backend/internal/modules/auth/dto"
	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"

	output "github.com/Toppira-Official/backend/internal/shared/dto"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
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

// SignUpWithEmailPassword godoc
//
//	@Summary	login with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Param		body	body		input.LoginWithEmailPasswordInput	true	"Login Input"
//	@Success	200		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Router		/auth/login-with-user-password [post]
func (hl *LoginHandler) LoginWithEmailPassword(c *gin.Context) {
	ctx := c.Request.Context()

	var input input.LoginWithEmailPasswordInput
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

	c.JSON(http.StatusOK, output.HttpOutput{
		Data: map[string]any{
			"user":         user,
			"access_token": accessToken,
		}},
	)
}

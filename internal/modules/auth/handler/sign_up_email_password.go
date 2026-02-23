package handler

import (
	"net/http"
	"strconv"

	authDto "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler/dto"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	userInput "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	createUserUsecase   userUsecase.CreateUserUsecase
	hashPasswordUsecase authUsecase.HashPasswordUsecase
	generateJwtUsecase  authUsecase.GenerateJwtUsecase
}

func NewSignUpHandler(
	createUserUsecase userUsecase.CreateUserUsecase,
	hashPasswordUsecase authUsecase.HashPasswordUsecase,
	generateJwtUsecase authUsecase.GenerateJwtUsecase,
) *SignUpHandler {
	return &SignUpHandler{
		createUserUsecase:   createUserUsecase,
		hashPasswordUsecase: hashPasswordUsecase,
		generateJwtUsecase:  generateJwtUsecase,
	}
}

// SignUpWithEmailPassword godoc
//
//	@Summary	sign up with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.SignUpWithEmailPasswordInput	true	"Sign Up Input"
//	@Success	201		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	409		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	ctx := c.Request.Context()

	var input authDto.SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	usecaseInput := &userInput.CreateUserInput{
		Email:    input.Email,
		Password: &input.Password,
		IsActive: false,
	}
	savedUser, err := hl.createUserUsecase.Execute(ctx, usecaseInput)
	if err != nil {
		c.Error(err)
		return
	}

	userIDString := strconv.Itoa(int(savedUser.ID))
	accessToken, err := hl.generateJwtUsecase.Execute(ctx, userIDString)
	if err != nil {
		c.Error(err)
		return
	}

	savedUser.Password = nil

	c.JSON(http.StatusCreated, output.HttpOutput[authDto.AuthOutput]{
		Data: authDto.AuthOutput{
			User:        output.ToUserOutput(savedUser),
			AccessToken: accessToken,
		},
	})
}

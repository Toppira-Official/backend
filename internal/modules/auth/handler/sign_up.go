package handler

import (
	"net/http"
	"strconv"
	"strings"

	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"
	"github.com/Toppira-Official/backend/internal/shared/dto"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	createUserUsecase   userUsecase.CreateUserUsecase
	hashPasswordUsecase authUsecase.HashPasswordUsecase
	generateJwtUsecase  authUsecase.GenerateJwtUsecase
}

func NewSignUpHandler(createUserUsecase userUsecase.CreateUserUsecase,
	hashPasswordUsecase authUsecase.HashPasswordUsecase,
	generateJwtUsecase authUsecase.GenerateJwtUsecase) *SignUpHandler {
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
//	@Param		body	body		SignUpWithEmailPasswordInput	true	"Sign Up Input"
//	@Success	200		{object}	dto.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	ctx := c.Request.Context()

	var input SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	hashedPassword, err := hl.hashPasswordUsecase.Execute(ctx, []byte(input.Password))
	if err != nil {
		c.Error(err)
		return
	}

	input.Email = strings.ToLower(input.Email)
	input.Password = hashedPassword
	user, err := hl.createUserUsecase.Execute(ctx, input.MapUser())
	if err != nil {
		c.Error(err)
		return
	}

	userIDString := strconv.Itoa(int(user.ID))
	accessToken, err := hl.generateJwtUsecase.Execute(ctx, userIDString)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.HttpOutput{
		Data: map[string]any{
			"user":         user,
			"access_token": accessToken,
		}},
	)
}

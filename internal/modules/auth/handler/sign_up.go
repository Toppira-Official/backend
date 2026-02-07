package handler

import (
	"net/http"

	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"
	"github.com/Toppira-Official/backend/internal/shared/dto"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SignUpHandler struct {
	createUserUsecase userUsecase.CreateUserUsecase
	logger            *zap.Logger
}

func NewSignUpHandler(createUserUsecase userUsecase.CreateUserUsecase) *SignUpHandler {
	return &SignUpHandler{createUserUsecase: createUserUsecase}
}

// SignUpWithEmailPassword godoc
//
//	@Summary	sign up with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Param		body	body		SignUpWithEmailPasswordInput	true	"Sign Up Input"
//	@Success	200		{object}	dto.HttpOutputDto
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	var input SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	user, err := hl.createUserUsecase.Execute(c.Request.Context(), input.MapUser())
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, dto.HttpOutputDto{
		Data: map[string]any{
			"user": user,
		}},
	)
}

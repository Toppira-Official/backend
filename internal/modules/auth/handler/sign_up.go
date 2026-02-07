package handler

import (
	"errors"
	"net/http"

	"github.com/Toppira-Official/backend/internal/domain/dto"
	apperrros "github.com/Toppira-Official/backend/internal/domain/errors"
	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"
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
//	@Failure	400		{object}	apperrros.ClientError
//	@Failure	500		{object}	apperrros.ClientError
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	var input SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, apperrros.E(apperrros.ErrUserInvalidData, err).Client())
		return
	}

	user, err := hl.createUserUsecase.Execute(c.Request.Context(), input.MapUser())
	if err != nil {
		var appErr *apperrros.AppError

		if errors.As(err, &appErr) {
			hl.logger.Error("signup failed",
				zap.String("code", string(appErr.Code)),
				zap.Error(appErr.Err),
			)

			status := apperrros.HTTPStatus(appErr.Code)

			c.JSON(status, appErr.Client())
			return
		}

		hl.logger.Error("unhandled error in SignUp", zap.Error(err))
		c.JSON(http.StatusInternalServerError, apperrros.E(apperrros.ErrServerInternalError, err).Client())
		return
	}

	c.JSON(http.StatusOK, dto.HttpOutputDto{
		Data: map[string]any{
			"user": user,
		}},
	)
}

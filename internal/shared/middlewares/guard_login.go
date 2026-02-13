package middlewares

import (
	"strconv"
	"strings"

	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

func GuardLogin(
	verifyJwtUsecase authUsecase.VerifyJwtUsecase,
	findUserByIDUsecase userUsecase.FindUserByIDUsecase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthTokenNotProvided}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := verifyJwtUsecase.Execute(c, token)
		if err != nil {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(claims.Subject)
		if err != nil {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}
		user, err := findUserByIDUsecase.Execute(ctx, uint(userID))
		if err != nil {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

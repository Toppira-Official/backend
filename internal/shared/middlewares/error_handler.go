package middlewares

import (
	"errors"
	"net/http"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0].Err

		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			logger.Error("application error", zap.String("code", string(appErr.Code)), zap.Error(appErr.Err))
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			return
		}

		logger.Error("unhandled error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, apperrors.E(apperrors.ErrServerInternalError, err).Client())
	}
}

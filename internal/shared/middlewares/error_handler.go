package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger, es *elasticsearch.TypedClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0].Err

		logEntry := struct {
			Level     string    `json:"level"`
			Message   string    `json:"message"`
			Method    string    `json:"method"`
			Path      string    `json:"path"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Level:     "error",
			Message:   err.Error(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			CreatedAt: time.Now(),
		}

		_, elErr := es.Index("system-logs").
			Document(logEntry).
			Do(context.Background())
		if elErr != nil {
			logger.Warn("failed to save log to elasticsearch", zap.Error(elErr))
		}

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

package middlewares

import (
	"math"
	"strconv"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func TokenBucket(l *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := l.Reserve()
		if !res.OK() {
			appErr := &errors.AppError{ClientError: errors.ClientError{Code: errors.ErrTooManyRequests}}
			c.JSON(errors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		delay := res.Delay()
		if delay == 0 {
			c.Next()
			return
		}

		res.Cancel()
		c.Header("Retry-After", strconv.Itoa(int(math.Ceil(delay.Seconds()))))
		appErr := &errors.AppError{ClientError: errors.ClientError{Code: errors.ErrTooManyRequests}}
		c.JSON(errors.HTTPStatus(appErr.Code), appErr.Client())
		c.Abort()
	}
}

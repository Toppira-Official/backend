package configs

import (
	"time"

	"golang.org/x/time/rate"
)

func NewRateLimiter(envs Environments) *rate.Limiter {
	return rate.NewLimiter(rate.Every(5*time.Second), 10)
}

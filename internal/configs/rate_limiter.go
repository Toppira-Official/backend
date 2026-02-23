package configs

import (
	"strconv"

	"golang.org/x/time/rate"
)

func NewRateLimiter(envs Environments) *rate.Limiter {
	qps, err := strconv.Atoi(envs.RATE_LIMIT_QPS.String())
	if err != nil {
		panic(err)
	}
	burst, err := strconv.Atoi(envs.RATE_LIMIT_BURST.String())
	if err != nil {
		panic(err)
	}
	return rate.NewLimiter(rate.Limit(qps), burst)
}

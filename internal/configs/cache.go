package configs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func NewCache(lc fx.Lifecycle, envs Environments) *redis.Client {
	redisDB, err := strconv.Atoi(envs.REDIS_DB.String())
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", envs.REDIS_HOST.String(), envs.REDIS_PORT.String()),
		Password: envs.REDIS_PASSWORD.String(),
		DB:       redisDB,
	})

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return redisClient.Close()
		},
	})

	return redisClient
}

package queues

import (
	"fmt"
	"strconv"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(envs configs.Environments) *asynq.Server {
	addr := fmt.Sprintf("%s:%s", envs.REDIS_HOST.String(), envs.REDIS_PORT.String())

	db, err := strconv.Atoi(envs.REDIS_DB.String())
	if err != nil {
		panic(err)
	}
	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     addr,
			Password: envs.REDIS_PASSWORD.String(),
			DB:       db,
		},
		asynq.Config{
			Concurrency: 3,
			Queues: map[string]int{
				"critical": 10,
				"default":  1,
			},
		},
	)
}

func NewMux() *asynq.ServeMux {
	return asynq.NewServeMux()
}

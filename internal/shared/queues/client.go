package queues

import (
	"fmt"
	"strconv"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	"github.com/hibiken/asynq"
)

type Client struct {
	c *asynq.Client
}

func NewClient(envs configs.Environments) *Client {
	db, err := strconv.Atoi(envs.REDIS_DB.String())
	if err != nil {
		panic(err)
	}
	return &Client{
		c: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", envs.REDIS_HOST.String(), envs.REDIS_PORT.String()),
			Password: envs.REDIS_PASSWORD.String(),
			DB:       db,
		}),
	}
}

func (cl *Client) Close() error { return cl.c.Close() }

func (cl *Client) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return cl.c.Enqueue(task, opts...)
}

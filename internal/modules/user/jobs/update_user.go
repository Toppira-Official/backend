package jobs

import (
	"context"
	"encoding/json"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"

	"github.com/hibiken/asynq"
)

const TypeUpdateUser = "user:update"

type UpdateUserJob struct {
	uc usecase.UpdateUserUsecase
}

func NewUpdateUserJob(uc usecase.UpdateUserUsecase) *UpdateUserJob {
	return &UpdateUserJob{uc: uc}
}

func (j *UpdateUserJob) Process(ctx context.Context, t *asynq.Task) error {
	var p input.UpdateUserInput
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return asynq.SkipRetry
	}

	_, err := j.uc.Execute(ctx, &p)
	if err == nil {
		return nil
	}

	return err
}

func Register(mux *asynq.ServeMux, updateUser *UpdateUserJob) {
	mux.HandleFunc(TypeUpdateUser, updateUser.Process)
}

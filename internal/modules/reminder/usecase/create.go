package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase/input"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/constants"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
)

type CreateReminderUsecase interface {
	Execute(ctx context.Context, input *input.CreateReminderInput) (*entities.Reminder, error)
}

type createReminderUsecase struct {
	repo    *repositories.Query
	breaker *gobreaker.CircuitBreaker[struct{}]
}

const CreateReminderRetryTime = 30 * time.Second

func NewCreateReminderUsecase(repo *repositories.Query) CreateReminderUsecase {
	settings := gobreaker.Settings{
		Name:        "create_reminder_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     CreateReminderRetryTime,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		IsSuccessful: func(err error) bool {
			if err == nil {
				return true
			}
			return false
		},
	}

	return &createReminderUsecase{repo: repo, breaker: gobreaker.NewCircuitBreaker[struct{}](settings)}
}

func (uc *createReminderUsecase) Execute(ctx context.Context, input *input.CreateReminderInput) (*entities.Reminder, error) {
	reminder := &entities.Reminder{
		Title:         input.Title,
		Description:   input.Description,
		Priority:      input.Priority,
		ScheduledAt:   input.ScheduledAt,
		ReminderTimes: input.ReminderTimes,
		Status:        constants.Pending,
		UserID:        input.UserID,
	}

	_, err := uc.breaker.Execute(func() (struct{}, error) {
		return struct{}{}, uc.repo.Reminder.WithContext(ctx).Create(reminder)
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return reminder, nil
}

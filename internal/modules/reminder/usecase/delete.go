package usecase

import (
	"context"
	"errors"
	"time"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
)

type DeleteReminderUsecase interface {
	Execute(ctx context.Context, id uint) error
}

type deleteReminderUsecase struct {
	repo    *repositories.Query
	breaker *gobreaker.CircuitBreaker[struct{}]
}

const DeleteReminderRetryTime = 30 * time.Second

func NewDeleteeReminderUsecase(repo *repositories.Query) DeleteReminderUsecase {
	settings := gobreaker.Settings{
		Name:        "delete_reminder_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     DeleteReminderRetryTime,
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

	return &deleteReminderUsecase{repo: repo, breaker: gobreaker.NewCircuitBreaker[struct{}](settings)}
}

func (uc *deleteReminderUsecase) Execute(ctx context.Context, id uint) error {
	_, err := uc.breaker.Execute(func() (struct{}, error) {
		_, err := uc.repo.Reminder.
			WithContext(ctx).
			Where(uc.repo.Reminder.BaseID.Eq(id)).
			Delete()
		return struct{}{}, err
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, err)
		}
		return apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return nil
}

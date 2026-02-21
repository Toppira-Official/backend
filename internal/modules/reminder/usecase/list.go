package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
)

type ListRemindersUsecase interface {
	Execute(ctx context.Context, userID uint, page, limit int) ([]*entities.Reminder, error)
}

type listRemindersUsecase struct {
	repo    *repositories.Query
	breaker *gobreaker.CircuitBreaker[[]*entities.Reminder]
}

const ListRemindersRetryTime = 30 * time.Second

func NewListRemindersUsecase(repo *repositories.Query) ListRemindersUsecase {
	settings := gobreaker.Settings{
		Name:        "list_reminders_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     ListRemindersRetryTime,
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

	return &listRemindersUsecase{repo: repo, breaker: gobreaker.NewCircuitBreaker[[]*entities.Reminder](settings)}
}

func (uc *listRemindersUsecase) Execute(ctx context.Context, userID uint, page, limit int) ([]*entities.Reminder, error) {
	offset := (page - 1) * limit
	reminders, err := uc.breaker.Execute(func() ([]*entities.Reminder, error) {
		return uc.repo.
			WithContext(ctx).Reminder.
			Where(uc.repo.Reminder.UserID.Eq(userID)).
			Order(uc.repo.Reminder.ScheduledAt.Desc()).
			Select(
				uc.repo.Reminder.BaseID,
				uc.repo.Reminder.Title,
				uc.repo.Reminder.ScheduledAt,
				uc.repo.Reminder.Status,
				uc.repo.Reminder.Priority,
			).
			Limit(limit).
			Offset(offset).
			Find()
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return reminders, nil
}

package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const updateUserRetryTime = 30 * time.Second

type UpdateUserUsecase interface {
	Execute(ctx context.Context, input *input.UpdateUserInput) (*entities.User, error)
}

type updateUserUsecase struct {
	repo         *repositories.Query
	hashPassword authUsecase.HashPasswordUsecase
	breaker      *gobreaker.CircuitBreaker[gen.ResultInfo]
}

func NewUpdateUserUsecase(repo *repositories.Query, hashPassword authUsecase.HashPasswordUsecase) UpdateUserUsecase {
	settings := gobreaker.Settings{
		Name:        "update_user_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     updateUserRetryTime,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		IsSuccessful: func(err error) bool {
			if err == nil {
				return true
			}
			if errors.Is(err, gorm.ErrDuplicatedKey) || errors.Is(err, gorm.ErrRecordNotFound) {
				return true
			}
			return false
		},
	}

	return &updateUserUsecase{repo: repo, hashPassword: hashPassword, breaker: gobreaker.NewCircuitBreaker[gen.ResultInfo](settings)}
}

func (uc *updateUserUsecase) Execute(ctx context.Context, input *input.UpdateUserInput) (*entities.User, error) {
	updateData := map[string]any{}

	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		updateData["name"] = strings.ToLower(name)
	}

	if input.Phone != nil {
		updateData["phone"] = *input.Phone
	}

	if input.Password != nil {
		hashed, err := uc.hashPassword.Execute(ctx, []byte(*input.Password))
		if err != nil {
			return nil, apperrors.E(apperrors.ErrServerInternalError, err)
		}
		updateData["password"] = hashed
	}

	if len(updateData) == 0 {
		return nil, apperrors.E(apperrors.ErrUserInvalidData, nil)
	}

	res, err := uc.breaker.Execute(func() (gen.ResultInfo, error) {
		return uc.repo.User.WithContext(ctx).
			Where(uc.repo.User.BaseID.Eq(input.ID)).
			Updates(updateData)
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, err)
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperrors.E(apperrors.ErrUserAlreadyExists, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	if res.RowsAffected == 0 {
		return nil, apperrors.E(apperrors.ErrUserNotFound, nil)
	}

	updatedUser, err := uc.repo.User.WithContext(ctx).Where(uc.repo.User.BaseID.Eq(input.ID)).First()
	if err != nil {
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return updatedUser, nil
}

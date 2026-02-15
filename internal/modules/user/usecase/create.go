package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
	"gorm.io/gorm"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, input *input.CreateUserInput) (*entities.User, error)
}

type createUserUsecase struct {
	repo         *repositories.Query
	hashPassword authUsecase.HashPasswordUsecase
	breaker      *gobreaker.CircuitBreaker[struct{}]
}

func NewCreateUserUsecase(repo *repositories.Query, hashPassword authUsecase.HashPasswordUsecase) CreateUserUsecase {
	settings := gobreaker.Settings{
		Name:        "create_user_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     configs.RetryDelay,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		IsSuccessful: func(err error) bool {
			if err == nil {
				return true
			}
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return true
			}
			return false
		},
	}

	return &createUserUsecase{repo: repo, hashPassword: hashPassword, breaker: gobreaker.NewCircuitBreaker[struct{}](settings)}
}

func (uc *createUserUsecase) Execute(ctx context.Context, input *input.CreateUserInput) (*entities.User, error) {
	if input.Email == "" {
		return nil, apperrors.E(
			apperrors.ErrUserInvalidData,
			errors.New("email is required"),
		)
	}

	email := strings.ToLower(input.Email)

	passwordHash, err := uc.hashPassword.Execute(ctx, []byte(*input.Password))
	if err != nil {
		return nil, apperrors.E(apperrors.ErrServerInternalError, err)
	}

	user := &entities.User{
		Email:          email,
		Phone:          input.Phone,
		Name:           input.Name,
		ProfilePicture: input.ProfilePicture,
		Password:       &passwordHash,
		IsActive:       false,
	}

	_, err = uc.breaker.Execute(func() (struct{}, error) {
		return struct{}{}, uc.repo.User.WithContext(ctx).Create(user)
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

	return user, nil
}

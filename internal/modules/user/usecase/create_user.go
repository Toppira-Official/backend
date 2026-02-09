package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/Toppira-Official/backend/internal/shared/entities"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/Toppira-Official/backend/internal/shared/repositories"
	"gorm.io/gorm"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, input *entities.User) (*entities.User, error)
}

type createUserUsecase struct {
	repo *repositories.Query
}

func NewCreateUserUsecase(repo *repositories.Query) CreateUserUsecase {
	return &createUserUsecase{repo: repo}
}

func (uc *createUserUsecase) Execute(ctx context.Context, input *entities.User) (*entities.User, error) {
	user := input
	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)
	}
	err := uc.repo.User.WithContext(ctx).Save(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperrors.E(apperrors.ErrUserAlreadyExists, err)
		}

		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return user, nil
}

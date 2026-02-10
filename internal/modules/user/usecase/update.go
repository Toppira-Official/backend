package usecase

import (
	"context"
	"errors"
	"strings"

	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	"github.com/Toppira-Official/backend/internal/shared/entities"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/Toppira-Official/backend/internal/shared/repositories"
	"gorm.io/gorm"
)

type UpdateUserUsecase interface {
	Execute(ctx context.Context, input *entities.User) (*entities.User, error)
}

type updateUserUsecase struct {
	repo         *repositories.Query
	hashPassword authUsecase.HashPasswordUsecase
}

func NewUpdateUserUsecase(repo *repositories.Query, hashPassword authUsecase.HashPasswordUsecase) UpdateUserUsecase {
	return &updateUserUsecase{repo: repo, hashPassword: hashPassword}
}

func (uc *updateUserUsecase) Execute(ctx context.Context, input *entities.User) (*entities.User, error) {
	user := input
	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)
	}
	if user.Password != nil {
		password, err := uc.hashPassword.Execute(ctx, []byte(*user.Password))
		if err != nil {
			return nil, apperrors.E(apperrors.ErrServerInternalError, err)
		}
		user.Password = &password
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

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

type FindUserByEmailUsecase interface {
	Execute(ctx context.Context, input string) (*entities.User, error)
}

type findUserByEmailUsecase struct {
	repo *repositories.Query
}

func NewFindUserByEmailUsecase(repo *repositories.Query) FindUserByEmailUsecase {
	return &findUserByEmailUsecase{repo: repo}
}

func (uc *findUserByEmailUsecase) Execute(ctx context.Context, input string) (*entities.User, error) {
	email := strings.ToLower(input)
	user, err := uc.repo.User.WithContext(ctx).Where(uc.repo.User.Email.Eq(email)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.E(apperrors.ErrUserNotFound, err)
		}

		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return user, nil
}

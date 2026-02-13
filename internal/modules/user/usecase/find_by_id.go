package usecase

import (
	"context"
	"errors"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"gorm.io/gorm"
)

type FindUserByIDUsecase interface {
	Execute(ctx context.Context, input uint) (*entities.User, error)
}

type findUserByIDUsecase struct {
	repo *repositories.Query
}

func NewFindUserByIDUsecase(repo *repositories.Query) FindUserByIDUsecase {
	return &findUserByIDUsecase{repo: repo}
}

func (uc *findUserByIDUsecase) Execute(ctx context.Context, id uint) (*entities.User, error) {
	user, err := uc.repo.User.WithContext(ctx).Where(uc.repo.User.BaseID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.E(apperrors.ErrUserNotFound, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return user, nil
}

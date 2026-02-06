package usecase

import (
	"context"

	"github.com/Toppira-Official/backend/internal/domain/repositories"
	"go.uber.org/zap"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
}

type createUserUsecase struct {
	repo   *repositories.Query
	logger *zap.Logger
}

func NewCreateUserUsecase(repo *repositories.Query, logger *zap.Logger) CreateUserUsecase {
	return &createUserUsecase{repo: repo, logger: logger}
}

func (uc *createUserUsecase) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	user := input.Map()
	err := uc.repo.User.WithContext(ctx).Save(user)
	if err != nil {
		uc.logger.Error("failed to create user", zap.Error(err))
		return CreateUserOutput{}, err
	}

	output := CreateUserOutput{
		ID:             user.ID,
		Email:          user.Email,
		Phone:          user.Phone,
		Name:           user.Name,
		ProfilePicture: user.ProfilePicture,
	}

	return output, nil
}

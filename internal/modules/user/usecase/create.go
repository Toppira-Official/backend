package usecase

import (
	"context"
	"errors"
	"strings"

	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	"github.com/Toppira-Official/backend/internal/modules/user/usecase/input"
	"github.com/Toppira-Official/backend/internal/shared/entities"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/Toppira-Official/backend/internal/shared/repositories"
	"gorm.io/gorm"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, input *input.CreateUserInput) (*entities.User, error)
}

type createUserUsecase struct {
	repo         *repositories.Query
	hashPassword authUsecase.HashPasswordUsecase
}

func NewCreateUserUsecase(repo *repositories.Query, hashPassword authUsecase.HashPasswordUsecase) CreateUserUsecase {
	return &createUserUsecase{repo: repo, hashPassword: hashPassword}
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

	err = uc.repo.User.WithContext(ctx).Create(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperrors.E(apperrors.ErrUserAlreadyExists, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return user, nil
}

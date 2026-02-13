package usecase

import (
	"context"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"golang.org/x/crypto/bcrypt"
)

type HashPasswordUsecase interface {
	Execute(ctx context.Context, password []byte) (string, error)
}

type hashPasswordUsecase struct{}

func NewCreateUserUsecase() HashPasswordUsecase {
	return &hashPasswordUsecase{}
}

func (uc *hashPasswordUsecase) Execute(ctx context.Context, password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return "", apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return string(bytes), nil
}

package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type VerifyPasswordUsecase interface {
	Execute(ctx context.Context, password, hash []byte) bool
}

type verifyPasswordUsecase struct{}

func NewVerifyPasswordUsecase() VerifyPasswordUsecase {
	return &verifyPasswordUsecase{}
}

func (uc *verifyPasswordUsecase) Execute(ctx context.Context, password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), password)
	return err == nil
}

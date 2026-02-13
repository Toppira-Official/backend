package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/golang-jwt/jwt/v5"
)

type GenerateJwtUsecase interface {
	Execute(ctx context.Context, id string) (string, error)
}

type generateJwtUsecase struct {
	envs configs.Environments
}

func NewGenerateJwtUsecase(envs configs.Environments) GenerateJwtUsecase {
	return &generateJwtUsecase{envs: envs}
}

func (uc *generateJwtUsecase) Execute(ctx context.Context, id string) (string, error) {
	expireTimeInHour, err := strconv.Atoi(uc.envs.JWT_EXPIRES_IN_HOURS.String())
	if err != nil {
		return "", apperrors.E(apperrors.ErrServerInternalError, err)
	}

	claims := &jwt.RegisteredClaims{
		Subject:   id,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireTimeInHour))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(uc.envs.JWT_SECRET.String()))
	if err != nil {
		return "", apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return tokenString, nil
}

package usecase

import (
	"context"
	"fmt"

	"github.com/Toppira-Official/backend/internal/configs"
	"github.com/golang-jwt/jwt/v5"
)

type VerifyJwtUsecase interface {
	Execute(ctx context.Context, tokenString string) (*jwt.RegisteredClaims, error)
}

type verifyJwtUsecase struct {
	envs configs.Environments
}

func NewVerifyJwtUsecase(envs configs.Environments) VerifyJwtUsecase {
	return &verifyJwtUsecase{envs: envs}
}

func (uc *verifyJwtUsecase) Execute(ctx context.Context, tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(uc.envs.JWT_SECRET.String()), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

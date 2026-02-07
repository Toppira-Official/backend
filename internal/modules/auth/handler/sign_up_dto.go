package handler

import "github.com/Toppira-Official/backend/internal/domain/entities"

type SignUpWithEmailPasswordInput struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

func (in *SignUpWithEmailPasswordInput) MapUser() *entities.User {
	return &entities.User{
		Email:    in.Email,
		Password: &in.Password,
	}
}

package handler

import "github.com/Toppira-Official/backend/internal/domain/entities"

type SignUpWithEmailPasswordInput struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

type SignUpWithEmailPasswordOutput struct {
	Message string         `json:"message"`
	Data    map[string]any `json:"data,omitempty"`
}

func (in *SignUpWithEmailPasswordInput) MapUser() *entities.User {
	return &entities.User{
		Email:    in.Email,
		Password: &in.Password,
	}
}

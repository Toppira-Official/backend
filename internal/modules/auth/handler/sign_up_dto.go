package handler

import "github.com/Toppira-Official/backend/internal/shared/entities"

type SignUpWithEmailPasswordInput struct {
	Email    string `binding:"required,email" json:"email" example:"user@example.com"`
	Password string `binding:"required,min=8,max=72" json:"password" example:"StrongPassword1234"`
} // @name SignUpWithEmailPasswordInput

func (in *SignUpWithEmailPasswordInput) MapUser() *entities.User {
	return &entities.User{
		Email:    in.Email,
		Password: &in.Password,
	}
}

package usecase

import "github.com/Toppira-Official/backend/internal/domain/entities"

type CreateUserInput struct {
	Email          string
	Phone          *string
	Name           *string
	ProfilePicture *string
	Password       *string
}

type CreateUserOutput struct {
	ID             uint
	Email          string
	Phone          *string
	Name           *string
	ProfilePicture *string
}

func (d *CreateUserInput) Map() *entities.User {
	return &entities.User{
		Email:          d.Email,
		Phone:          d.Phone,
		Name:           d.Name,
		ProfilePicture: d.ProfilePicture,
		Password:       d.Password,
	}
}

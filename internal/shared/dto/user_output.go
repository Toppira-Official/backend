package dto

import (
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
)

type UserOutput struct {
	ID             uint      `json:"id"`
	Email          string    `json:"email"`
	Phone          *string   `json:"phone,omitempty"`
	IsActive       bool      `json:"is_active"`
	Name           *string   `json:"name,omitempty"`
	ProfilePicture *string   `json:"profile_picture,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
} //	@name	UserOutput

func ToUserOutput(u *entities.User) *UserOutput {
	if u == nil {
		return nil
	}
	return &UserOutput{
		ID:             u.ID,
		Email:          u.Email,
		Phone:          u.Phone,
		IsActive:       u.IsActive,
		Name:           u.Name,
		ProfilePicture: u.ProfilePicture,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

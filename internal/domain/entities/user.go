package entities

import "time"

type User struct {
	ID             uint      `json:"id"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           *string   `json:"name,omitempty"`
	Phone          *string   `json:"phone,omitempty"`
	Password       *string   `json:"-"`
	ProfilePicture *string   `json:"profile_picture,omitempty"`
}

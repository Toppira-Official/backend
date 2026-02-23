package dto

import sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"

type AuthOutput struct {
	User        *sharedDto.UserOutput `json:"user"`
	AccessToken string                `json:"access_token"`
} //	@name	AuthOutput

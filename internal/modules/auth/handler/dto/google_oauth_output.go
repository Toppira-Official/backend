package dto

import "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"

type GoogleOAuthOutput struct {
	User         *usecase.GoogleUserInfo `json:"user"`
	AccessToken  string                 `json:"access_token"`
} //	@name	GoogleOAuthOutput

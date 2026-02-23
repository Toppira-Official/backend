package dto

import output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"

type UpdateMeOutput struct {
	User *output.UserOutput `json:"user"`
} //	@name	UpdateMeOutput

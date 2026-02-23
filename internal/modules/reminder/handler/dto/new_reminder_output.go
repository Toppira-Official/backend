package dto

import "github.com/Toppira-Official/Reminder_Server/internal/shared/entities"

type NewReminderOutput struct {
	Reminder *entities.Reminder `json:"reminder"`
} //	@name	NewReminderOutput

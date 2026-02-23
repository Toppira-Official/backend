package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	reminderInput "github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase/input"
	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type NewReminderHandler struct {
	createReminderUsecase usecase.CreateReminderUsecase
}

func NewNewReminderHandler(
	createReminderUsecase usecase.CreateReminderUsecase,
) *NewReminderHandler {
	return &NewReminderHandler{
		createReminderUsecase: createReminderUsecase,
	}
}

// NewReminder godoc
//
//	@Summary	creates new reminder
//	@Tags		Reminder
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.NewReminderInput	true	"New Reminder Input"
//	@Success	201		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Security	BearerAuth
//	@Router		/reminder [post]
func (hl *NewReminderHandler) NewReminder(c *gin.Context) {
	ctx := c.Request.Context()

	userVal, exists := c.Get("user")
	if !exists {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	user, ok := userVal.(*entities.User)
	if !ok {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	var input dto.NewReminderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrReminderInvalidData, err))
		return
	}

	usecaseInput := &reminderInput.CreateReminderInput{
		Title:         input.Title,
		Description:   input.Description,
		Priority:      input.Priority,
		ReminderTimes: input.ReminderTimes,
		ScheduledAt:   input.ScheduledAt,
		UserID:        user.ID,
	}
	newReminder, err := hl.createReminderUsecase.Execute(ctx, usecaseInput)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, output.HttpOutput[dto.NewReminderOutput]{
		Data: dto.NewReminderOutput{
			Reminder: newReminder,
		},
	})
}

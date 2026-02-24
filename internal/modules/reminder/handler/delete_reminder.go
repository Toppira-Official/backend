package handler

import (
	"net/http"

	reminderInput "github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/dto/input"
	reminderOutput "github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/dto/output"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type DeleteRemindersHandler struct {
	deleteReminderUsecase usecase.DeleteReminderUsecase
}

func NewDeleteRemindersHandler(deleteReminderUsecase usecase.DeleteReminderUsecase) *DeleteRemindersHandler {
	return &DeleteRemindersHandler{deleteReminderUsecase: deleteReminderUsecase}
}

// DeleteReminder godoc
//
//	@Summary	deletes the given reminder
//	@Tags		Reminder
//	@Produce	json
//	@Param		id	path		int	true	"Reminder ID"
//	@Success	200	{object}	sharedDto.HttpOutput[reminderOutput.DeleteReminderOutput]
//	@Failure	500	{object}	apperrors.ClientError
//	@Failure	503	{object}	apperrors.ClientError
//	@Security	BearerAuth
//	@Router		/reminder/{id} [delete]
func (hl *DeleteRemindersHandler) DeleteReminder(c *gin.Context) {
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

	var q reminderInput.DeleteReminderInput
	if err := c.ShouldBindUri(&q); err != nil {
		c.Error(apperrors.E(apperrors.ErrReminderInvalidData))
		return
	}

	if err := hl.deleteReminderUsecase.Execute(c.Request.Context(), q.ID, user.ID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, sharedDto.HttpOutput[reminderOutput.DeleteReminderOutput]{
		Data: reminderOutput.DeleteReminderOutput{
			ID: q.ID,
		},
	})
}

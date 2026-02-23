package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"

	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"

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
//	@Success	200	{object}	output.HttpOutput
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

	var q dto.DeleteReminderInput
	if err := c.ShouldBindUri(&q); err != nil {
		c.Error(apperrors.E(apperrors.ErrReminderInvalidData))
		return
	}

	if err := hl.deleteReminderUsecase.Execute(c.Request.Context(), q.ID, user.ID); err != nil {
		c.Error(apperrors.E(apperrors.ErrReminderInvalidData))
		return
	}

	c.JSON(http.StatusOK, output.HttpOutput[dto.DeleteReminderOutput]{
		Data: dto.DeleteReminderOutput{
			ID: q.ID,
		},
	})
}

package handler

import (
	"net/http"

	output "github.com/Toppira-Official/backend/internal/shared/dto"
	"github.com/Toppira-Official/backend/internal/shared/entities"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type GetMeHandler struct{}

func NewGetMeHandler() *GetMeHandler {
	return &GetMeHandler{}
}

// GetMe godoc
//
//	@Summary	get my(user) data
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	output.HttpOutput
//	@Failure	400	{object}	apperrors.ClientError
//	@Failure	401	{object}	apperrors.ClientError
//	@Failure	500	{object}	apperrors.ClientError
//	@Failure	503	{object}	apperrors.ClientError
//	@Router		/user/me [get]
func (hl *GetMeHandler) GetMyInfo(c *gin.Context) {
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

	c.JSON(http.StatusOK, output.HttpOutput{
		Data: map[string]any{
			"user": user,
		},
	})
}

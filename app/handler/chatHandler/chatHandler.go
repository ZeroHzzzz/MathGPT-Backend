package chathandler

import (
	"mathgpt/app/apiException"
	chatservices "mathgpt/app/services/chatServices"
	"mathgpt/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

type newChatRequest struct {
	UserID string `json:"user_id"`
}

func NewChatHandler(c *gin.Context) {
	var req newChatRequest
	currentID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	if req.UserID != currentID {
		logger.Default.Error(c, "User %s tried to access user %s's profile", currentID, req.UserID)
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	chat_id, err := chatservices.NewChat(req.UserID)
	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"chat_id": chat_id,
	})
}

package chathandler

import (
	"log"
	"mathgpt/app/apiException"
	chatservices "mathgpt/app/services/chatServices"
	messageservices "mathgpt/app/services/messageServices"
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

// 获取会话列表
type getChatRequest struct {
	UserID   string `json:"user_id"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Search   string `json:"search"`
}

func GetChatHandler(c *gin.Context) {
	var req getChatRequest
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

	chats, err := chatservices.GetChatList(req.UserID, req.Search, req.Page, req.PageSize)
	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"chats": chats,
	})
}

func DelChatHandler(c *gin.Context) {
	currentID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	chatID := c.Param("chat_id")
	if chatID == "" {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	err := chatservices.DelChat(currentID.(string), chatID)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	err = messageservices.DelMessage(chatID)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

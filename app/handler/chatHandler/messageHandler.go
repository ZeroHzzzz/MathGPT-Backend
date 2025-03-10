package chathandler

import (
	"mathgpt/app/apiException"
	messageservices "mathgpt/app/services/messageServices"

	"github.com/gin-gonic/gin"
)

type newMessage struct {
	ChatID  string `json:"chat_id"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessage(c *gin.Context) {
	var req newMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request"})
		return
	}

	_, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	err := messageservices.CreateMessage(req.ChatID, req.Role, req.Content)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}

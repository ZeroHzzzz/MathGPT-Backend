package chathandler

import (
	"mathgpt/app/apiException"
	llmservices "mathgpt/app/services/llmServices"
	messageservices "mathgpt/app/services/messageServices"
	"mathgpt/app/utils"

	"github.com/gin-gonic/gin"
)

type newQuestion struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewQuestion(c *gin.Context) {
	var req newQuestion
	chatID := c.Param("chatID")
	if err := c.ShouldBindJSON(&req); err != nil || chatID == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request"})
		return
	}

	_, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	err := messageservices.CreateMessage(chatID, req.Role, req.Content)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	ans, err := llmservices.GetAnswer(req.Content)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"role":    "llm",
		"content": ans,
	})
}

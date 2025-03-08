package userhandler

import (
	"mathgpt/app/apiException"
	userservices "mathgpt/app/services/userServices"
	"mathgpt/app/utils"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func RegisterByEmailHandler(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, apiException.ParamError)
		return
	}

	user, err := userservices.CreateUser(req.Email, req.Phone, req.Password)
	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, user)
}

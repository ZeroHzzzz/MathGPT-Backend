package userhandler

import (
	"mathgpt/app/apiException"
	userservices "mathgpt/app/services/userServices"
	"mathgpt/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type resetPassRequest struct {
	UserID      string `json:"user_id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func ResetPassHandler(c *gin.Context) {
	var req resetPassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	// 验证用户密码
	_, err := userservices.GetUserByIDAndPass(req.UserID, req.Password)
	if err == gorm.ErrRecordNotFound {
		c.AbortWithError(400, apiException.NoThatUserOrPasswordWrong)
		return
	}

	err = userservices.UpdateUser(req.UserID, map[string]interface{}{
		"password": req.NewPassword,
	})

	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type updateProfileRequest struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func UpdateProfileHandler(c *gin.Context) {
	var req updateProfileRequest
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

	err := userservices.UpdateUser(req.UserID, map[string]interface{}{
		"username": req.Username,
	})
	if err != nil {
		if err == apiException.UserNotFind {
			c.AbortWithError(400, err)
			return
		} else {
			c.AbortWithError(500, apiException.ServerError)
			return
		}
	}

	user, err := userservices.GetUserByID(req.UserID)
	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, user)
}

func GetUserProfileHandler(c *gin.Context) {
	currentID, exists := c.Get("user_id")
	if !exists {
		logger.Default.Error(c, "User not login")
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	userID := c.Param("userID")
	if userID == "" {
		logger.Default.Error(c, "no param userID")
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	if userID != currentID {
		logger.Default.Error(c, "User %s tried to access user %s's profile", currentID, userID)
		c.AbortWithError(400, apiException.NotLogin)
		return
	}

	user, err := userservices.GetUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(400, apiException.UserNotFind)
			return
		} else {
			c.AbortWithError(500, apiException.ServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, user)
}

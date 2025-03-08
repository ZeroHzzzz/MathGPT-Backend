package userhandler

import (
	"mathgpt/app/apiException"
	"mathgpt/app/midwares"
	userservices "mathgpt/app/services/userServices"
	"mathgpt/app/utils"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Password string `json:"password"`
	Account  string `json:"account"`
}

func LoginByIDHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	// account, err := strconv.ParseUint(req.Account, 0, 0)
	// if err != nil {
	// 	c.AbortWithError(400, apiException.ParamError)
	// 	return
	// }

	user, err := userservices.GetUserByIDAndPass(req.Account, req.Password)
	if err != nil && user.ID == req.Account {
		c.AbortWithError(400, apiException.NoThatUserOrPasswordWrong)
		return
	}

	token, err := midwares.CreateJWT(user.ID)

	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	c.Set("user_id", user.ID)

	utils.JsonSuccessResponse(c, gin.H{
		"user":      user,
		"token":     token,
		"expiresIn": midwares.Duration,
	})
}

func LoginByEmailHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	user, err := userservices.GetUserByEmailAndPass(req.Account, req.Password)
	if err != nil {
		c.AbortWithError(400, apiException.NoThatUserOrPasswordWrong)
		return
	}

	token, err := midwares.CreateJWT(user.ID)

	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user":  user,
		"token": token,
	})
}

func LoginByPhoneHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	user, err := userservices.GetUserByPhoneAndPass(req.Account, req.Password)
	if err != nil {
		c.AbortWithError(400, apiException.NoThatUserOrPasswordWrong)
		return
	}

	token, err := midwares.CreateJWT(user.ID)

	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	c.Set("user_id", user.ID)

	utils.JsonSuccessResponse(c, gin.H{
		"user":      user,
		"token":     token,
		"expiresIn": midwares.Duration,
	})
}

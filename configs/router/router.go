package router

import (
	userhandler "mathgpt/app/handler/userHandler"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"
	api := r.Group(pre)
	{
		api.POST("/login", userhandler.LoginByIDHandler)
	}
}

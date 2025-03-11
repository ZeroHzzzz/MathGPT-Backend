package router

import (
	chathandler "mathgpt/app/handler/chatHandler"
	userhandler "mathgpt/app/handler/userHandler"
	"mathgpt/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"
	api := r.Group(pre)
	{

	}
	auth := api.Group("/auth")
	{
		auth.POST("/login", userhandler.LoginByIDHandler)
	}
	user := api.Group("/user")
	{
		user.Use(midwares.JWTMiddleware())
		user.PATCH("/reset_pass", userhandler.ResetPassHandler)
		user.GET("/:userID", userhandler.GetUserProfileHandler)
		user.PUT("/update", userhandler.UpdateProfileHandler)
	}
	chat := api.Group("/chat")
	{
		chat.Use(midwares.JWTMiddleware())
		chat.POST("/new", chathandler.NewChatHandler)
	}
	message := chat.Group("/message")
	{
		message.GET("/question/:chatID", chathandler.NewQuestion)
	}
}

package main

import (
	"log"
	"mathgpt/app/midwares"
	"mathgpt/configs/config"
	"mathgpt/configs/database/mongodb"
	"mathgpt/configs/database/mysql"
	"mathgpt/configs/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port = ":" + config.Config.GetString("server.port")

func main() {

	mongodb.Init()
	mysql.Init()

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)

	router.Init(r)
	err := r.Run(port)
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}

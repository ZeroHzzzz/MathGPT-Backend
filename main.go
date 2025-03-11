package main

import (
	"log"
	"mathgpt/app/midwares"
	"mathgpt/configs/config"
	"mathgpt/configs/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port = ":" + config.Config.GetString("server.port")

func main() {

	// mongodb.Init()
	// mysql.Init()

	// ctx := context.Background()
	// testData := bson.M{
	// 	"message":    "Hello, this is a test message",
	// 	"created_at": time.Now(),
	// }
	// col := mongodb.MDB.Collection("chats")
	// col.InsertOne(ctx, testData)

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

	// chatID := "67d08ea411f471a135098bdb"

	// // 尝试转换为 ObjectID
	// objID, err := primitive.ObjectIDFromHex(chatID)
	// if err != nil {
	// 	log.Printf("the provided hex string is not a valid ObjectID: %v", err)
	// 	return
	// }

	// fmt.Println("Successfully converted to ObjectID:", objID.Hex())
}

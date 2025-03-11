package mongodb

import (
	"context"
	"fmt"
	"log"
	"mathgpt/configs/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type MongoDB struct {
// 	Client *mongo.Client
// 	DB     *mongo.Database
// }

var MDB *mongo.Database

// var MongoDB *mongo.Collection

func init() {
	user := config.Config.GetString("mongodb.user")
	pass := config.Config.GetString("mongodb.pass")
	host := config.Config.GetString("mongodb.host")
	port := config.Config.GetString("mongodb.port")
	db := config.Config.GetString("mongodb.db")
	// collection := config.Config.GetString("mongodb.collection")

	dsn := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", user, pass, host, port, db)
	// 构建 MongoDB 连接字符串
	if user == "" || pass == "" {
		dsn = fmt.Sprintf("mongodb://%v:%v/%v", host, port, db)
	}

	// 使用 dsn 连接 MongoDB
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:" + err.Error())
		return
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Failed to ping MongoDB:" + err.Error())
		return
	}

	// Print a log message to indicate successful connection to MongoDB
	log.Println("Connected to MongoDB")

	MDB = client.Database(db)
}

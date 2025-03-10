package messageservices

import (
	"context"
	"mathgpt/configs/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.Database.MDB.DB.Collection("messages")

func CreateMessage(chatID, Role, Content string) error {
	message := bson.M{
		"chat_id":    chatID,
		"role":       Role,
		"content":    Content,
		"created_at": time.Now(),
	}

	_, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}

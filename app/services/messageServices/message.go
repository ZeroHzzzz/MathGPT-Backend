package messageservices

import (
	"context"
	"mathgpt/configs/database/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = mongodb.MDB.Collection("messages")

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

package messageservices

import (
	"context"
	"mathgpt/app/models"
	"mathgpt/configs/database/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetMessage(chatID string, page, pageSize int64) ([]models.Message, error) {
	var messages []models.Message

	// Ensure page and pageSize are valid
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	skip := (page - 1) * pageSize

	// Create filter with chatID
	filter := bson.D{{Key: "chat_id", Value: chatID}}

	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}). // Sort by creation time descending
		SetSkip(skip).                                   // Skip documents for pagination
		SetLimit(pageSize)

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func DelMessage(chatID string) error {
	_, err := collection.DeleteMany(context.Background(), bson.D{{Key: "chat_id", Value: chatID}})
	if err != nil {
		return err
	}

	return nil
}

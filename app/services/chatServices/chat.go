package chatservices

import (
	"context"
	"mathgpt/app/models"
	"mathgpt/configs/database/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = mongodb.MDB.Collection("chats")

func NewChat(userID string) (primitive.ObjectID, error) {
	newChat := models.Chat{
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection.InsertOne(context.TODO(), newChat)

	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func GetChatList(userID string, search string, page, pageSize int64) ([]models.Chat, error) {
	var chats []models.Chat

	// Ensure page and pageSize are valid
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	skip := (page - 1) * pageSize

	// Create filter with userID and optional title search
	filter := bson.D{{Key: "user_id", Value: userID}}
	if search != "" {
		filter = append(filter, bson.E{
			Key: "title",
			Value: bson.D{{
				Key:   "$regex",
				Value: primitive.Regex{Pattern: search, Options: "i"},
			}},
		})
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "updated_at", Value: -1}}). // Sort by creation time descending
		SetSkip(skip).                                   // Skip documents for pagination
		SetLimit(pageSize)                               // Limit results per page

	cursor, err := collection.Find(context.TODO(),
		filter,
		opts,
	)

	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &chats)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func DelChat(userID, chatID string) error {
	objID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.D{
		{Key: "_id", Value: objID},
		{Key: "user_id", Value: userID},
	})

	return err
}

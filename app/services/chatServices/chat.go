package chatservices

import (
	"context"
	"mathgpt/app/models"
	"mathgpt/configs/database/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    primitive.ObjectID `bson:"chat_id"`
	Role      string             `bson:"role"` // "user" 或 "assistant"
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}

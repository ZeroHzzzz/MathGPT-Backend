package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SessionID primitive.ObjectID `bson:"session_id"`
	Role      string             `bson:"role"` // "user" 或 "assistant"
	Content   string             `bson:"content"`
	Timestamp time.Time          `bson:"timestamp"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatMessage struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Order     uint               `json:"order" bson:"order"`
	Text      string             `json:"text" bson:"text"`
	Type      string             `json:"type" bson:"type"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	CreateAt  time.Time          `json:"created_at" bson:"created_at"`
}

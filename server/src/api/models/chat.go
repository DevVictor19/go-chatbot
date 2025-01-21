package models

import (
	"time"
)

type Chat struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	BusinessId    string    `json:"business_id" bson:"business_id"`
	TotalMessages int       `json:"total_messages" bson:"total_messages"`
	Active        bool      `json:"active" bson:"active"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	CreateAt      time.Time `json:"created_at" bson:"created_at"`
}

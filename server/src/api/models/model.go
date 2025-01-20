package models

import "time"

type Model struct {
	ID        string    `json:"id" bson:"_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	CreateAt  time.Time `json:"created_at" bson:"created_at"`
}

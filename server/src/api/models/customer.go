package models

import (
	"time"
)

type Customer struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"-" bson:"password"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	CreateAt  time.Time `json:"created_at" bson:"created_at"`
}

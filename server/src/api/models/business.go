package models

import (
	"time"
)

type colorSchema struct {
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
	Paper     string `json:"paper" bson:"paper"`
	Text      string `json:"text" bson:"text"`
}

type Business struct {
	ID          string      `json:"id" bson:"_id,omitempty"`
	CustomerId  string      `json:"customer_id" bson:"customer_id"`
	Email       string      `json:"email" bson:"email"`
	Specialty   string      `json:"specialty" bson:"specialty"`
	History     string      `json:"history" bson:"history"`
	ColorSchema colorSchema `json:"color_schema" bson:"color_schema"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
	CreateAt    time.Time   `json:"created_at" bson:"created_at"`
}

package models

import (
	"time"
)

type Product struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	BusinessId    string    `json:"business_id" bson:"business_id"`
	PhotoURL      string    `json:"photo_url" bson:"photo_url"`
	Name          string    `json:"name" bson:"name"`
	Description   string    `json:"description" bson:"description"`
	StockQuantity int       `json:"stock_qnt" bson:"stock_qnt"`
	Price         float64   `json:"price" bson:"price"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	CreateAt      time.Time `json:"created_at" bson:"created_at"`
}

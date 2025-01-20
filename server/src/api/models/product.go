package models

type Product struct {
	Model
	BusinessId    string  `json:"business_id" bson:"business_id"`
	PhotoURL      string  `json:"photo_url" bson:"photo_url"`
	Name          string  `json:"name" bson:"name"`
	Description   string  `json:"description" bson:"description"`
	StockQuantity int     `json:"stock_qnt" bson:"stock_qnt"`
	Price         float64 `json:"price" bson:"price"`
}

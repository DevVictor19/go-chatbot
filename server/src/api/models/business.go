package models

type colorSchema struct {
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
	Paper     string `json:"paper" bson:"paper"`
	Text      string `json:"text" bson:"text"`
}

type Business struct {
	Model
	CustomerId  string      `json:"customer_id" bson:"customer_id"`
	Email       string      `json:"email" bson:"email"`
	Specialty   string      `json:"specialty" bson:"specialty"`
	History     string      `json:"history" bson:"history"`
	ColorSchema colorSchema `json:"color_schema" bson:"color_schema"`
}

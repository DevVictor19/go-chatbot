package models

type Chat struct {
	Model
	BusinessId    string `json:"business_id" bson:"business_id"`
	TotalMessages int    `json:"total_messages" bson:"total_messages"`
	Active        bool   `json:"active" bson:"active"`
}

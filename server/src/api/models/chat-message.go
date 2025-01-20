package models

type ChatMessage struct {
	Model
	Order uint   `json:"order" bson:"order"`
	Text  string `json:"text" bson:"text"`
	Type  string `json:"type" bson:"type"`
}

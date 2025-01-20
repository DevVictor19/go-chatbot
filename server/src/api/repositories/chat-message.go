package repositories

import (
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChatMessageRepository struct {
	*Repository[models.ChatMessage]
}

var chatMsgRepository *ChatMessageRepository

func NewChatMessageRepository(db *mongo.Database) *ChatMessageRepository {
	if chatMsgRepository == nil {
		chatMsgRepository = &ChatMessageRepository{
			Repository: newRepository[models.ChatMessage](db, "chat-message"),
		}
		return chatMsgRepository
	}

	return chatMsgRepository
}

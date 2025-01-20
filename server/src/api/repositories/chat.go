package repositories

import (
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChatRepository struct {
	*Repository[models.Chat]
}

var chatRepository *ChatRepository

func NewChatRepository(db *mongo.Database) *ChatRepository {
	if chatRepository == nil {
		chatRepository = &ChatRepository{
			Repository: newRepository[models.Chat](db, "chat"),
		}
		return chatRepository
	}

	return chatRepository
}

package repositories

import (
	"server/src/api"
	"server/src/api/models"
)

var ChatMessageRepository = newRepository[models.ChatMessage](api.GetDatabase(), "chat-message")

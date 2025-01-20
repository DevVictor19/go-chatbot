package repositories

import (
	"server/src/api"
	"server/src/api/models"
)

var ChatRepository = newRepository[models.Chat](api.GetDatabase(), "chat")

package repositories

import (
	"server/src/api"
	"server/src/api/models"
)

var BusinessRepository = newRepository[models.Business](api.GetDatabase(), "business")

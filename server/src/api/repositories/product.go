package repositories

import (
	"server/src/api"
	"server/src/api/models"
)

var ProductRepository = newRepository[models.Product](api.GetDatabase(), "product")

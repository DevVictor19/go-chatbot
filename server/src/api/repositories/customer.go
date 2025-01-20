package repositories

import (
	"server/src/api"
	"server/src/api/models"
)

var CustomerRepository = newRepository[models.Customer](api.GetDatabase(), "customer")

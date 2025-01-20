package repositories

import (
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BusinessRepository struct {
	*Repository[models.Business]
}

var businessRepository *BusinessRepository

func NewBusinessRepository(db *mongo.Database) *BusinessRepository {
	if businessRepository == nil {
		businessRepository = &BusinessRepository{
			Repository: newRepository[models.Business](db, "business"),
		}
		return businessRepository
	}

	return businessRepository
}

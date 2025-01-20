package repositories

import (
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository struct {
	*Repository[models.Product]
}

var productRepository *ProductRepository

func NewProductRepository(db *mongo.Database) *ProductRepository {
	if productRepository == nil {
		productRepository = &ProductRepository{
			Repository: newRepository[models.Product](db, "product"),
		}
		return productRepository
	}

	return productRepository
}

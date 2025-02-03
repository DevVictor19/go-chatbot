package repositories

import (
	"context"
	"fmt"
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ProductRepository struct {
	*Repository[models.Product]
}

func (r *ProductRepository) FindByIdAndBusinessId(
	ctx context.Context,
	productId,
	businessId string,
) (*models.Product, error) {
	var product models.Product

	filter := bson.D{
		{Key: "_id", Value: productId},
		{Key: "business_id", Value: businessId},
	}

	err := r.coll.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return &product, nil
}

func (r *ProductRepository) FindAllPaginatedByBusinessId(
	ctx context.Context,
	businessId string,
	page,
	limit int64) (*PaginatedResult[models.Product], error) {
	if page < 0 {
		page = 0
	}
	if limit == 0 {
		limit = 10
	}

	skip := page * limit

	filter := bson.D{{Key: "business_id", Value: businessId}}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	cur, err := r.coll.Find(ctx, filter,
		options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Product
	for cur.Next(ctx) {
		var result models.Product
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &PaginatedResult[models.Product]{
		Page:         page,
		ItemsPerPage: limit,
		Total:        total,
		Data:         &results,
	}, nil
}

func (r *ProductRepository) DeleteByIdAndBusinessId(
	ctx context.Context,
	productId,
	businessId string,
) error {
	filter := bson.D{
		{Key: "_id", Value: productId},
		{Key: "business_id", Value: businessId},
	}

	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
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

package repositories

import (
	"context"
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BusinessRepository struct {
	*Repository[models.Business]
}

func (r *BusinessRepository) FindAllPaginatedByCustomerId(
	ctx context.Context,
	customerId string,
	page,
	limit int64) (*PaginatedResult[models.Business], error) {
	if page < 0 {
		page = 0
	}
	if limit == 0 {
		limit = 10
	}

	skip := page * limit

	filter := bson.D{{Key: "customer_id", Value: customerId}}

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

	var results []*models.Business
	for cur.Next(ctx) {
		var result models.Business
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &PaginatedResult[models.Business]{
		Page:         page,
		ItemsPerPage: limit,
		Total:        total,
		Data:         results,
	}, nil
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

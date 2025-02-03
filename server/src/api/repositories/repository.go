package repositories

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PaginatedResult[T interface{}] struct {
	Page         int64 `json:"page"`
	ItemsPerPage int64 `json:"items_per_page"`
	Total        int64 `json:"total"`
	Data         *[]*T `json:"data"`
}

type Repository[T interface{}] struct {
	coll *mongo.Collection
}

func newRepository[T interface{}](db *mongo.Database, collection string) *Repository[T] {
	return &Repository[T]{coll: db.Collection(collection)}
}

func (r *Repository[T]) Create(ctx context.Context, model *T) error {
	v := reflect.ValueOf(model).Elem()
	field := v.FieldByName("ID")

	if field.IsValid() && field.CanSet() {
		if field.Kind() == reflect.String && field.String() == "" {
			field.SetString(primitive.NewObjectID().Hex())
		}
	}

	_, err := r.coll.InsertOne(ctx, model)

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (r *Repository[T]) FindById(ctx context.Context, id string) (*T, error) {
	var result T

	err := r.coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return &result, nil
}

func (r *Repository[T]) FindAll(ctx context.Context) ([]*T, error) {
	cur, err := r.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var results []*T

	for cur.Next(ctx) {
		var result T
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *Repository[T]) FindAllPaginated(ctx context.Context, page, limit int64) (*PaginatedResult[T], error) {
	if page < 0 {
		page = 0
	}
	if limit == 0 {
		limit = 10
	}

	skip := page * limit

	total, err := r.coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	cur, err := r.coll.Find(ctx, bson.D{},
		options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*T
	for cur.Next(ctx) {
		var result T
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &PaginatedResult[T]{
		Page:         page,
		ItemsPerPage: limit,
		Total:        total,
		Data:         &results,
	}, nil
}

func (r *Repository[T]) Update(ctx context.Context, model *T) error {
	idField := reflect.ValueOf(model).Elem().FieldByName("ID")
	if !idField.IsValid() || idField.Kind() != reflect.String {
		return fmt.Errorf("model does not have a valid 'ID' field")
	}

	id := idField.String()
	if id == "" {
		return fmt.Errorf("model ID cannot be empty")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": model}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *Repository[T]) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}

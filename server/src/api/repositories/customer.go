package repositories

import (
	"context"
	"fmt"
	"server/src/api/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CustomerRepository struct {
	*Repository[models.Customer]
}

func (r *CustomerRepository) FindByEmail(ctx context.Context, email string) (*models.Customer, error) {
	var customer models.Customer

	err := r.coll.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&customer)

	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return &customer, nil
}

var customerRepository *CustomerRepository

func NewCustomerRepository(db *mongo.Database) *CustomerRepository {
	if customerRepository == nil {
		customerRepository = &CustomerRepository{
			Repository: newRepository[models.Customer](db, "customer"),
		}
		return customerRepository
	}

	return customerRepository
}

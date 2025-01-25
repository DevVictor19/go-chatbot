package database

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func GetDatabase() *mongo.Database {
	if mongoDatabase == nil {
		mongoDatabase = mongoClient.Database("chatbot")
		return mongoDatabase
	}
	return mongoDatabase
}

func getConnectionUrl() string {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	return fmt.Sprintf("mongodb://%s:%s@%s:27017", user, pwd, host)
}

func InitDatabaseConnection() {
	client, err := mongo.Connect(options.Client().ApplyURI(getConnectionUrl()))

	if err != nil {
		log.Fatalf("Failed to create a new mongo client: %v", err)
	}

	mongoClient = client
}

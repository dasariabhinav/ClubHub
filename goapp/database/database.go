package database

import (
	"context"
	"log"
	"os"
	"project/initializers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {
	initializers.LoadEnvVariables()

}

func ConnectDB() {
	dbURI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(dbName).Collection(colName)
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
)

type MongoDBWrapper struct {
	client *mongo.Client
}

func ConMongoDatabase() DatabaseMongo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsnMongo := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=%s",
		os.Getenv("MONGO_DB_USERNAME"),
		os.Getenv("MONGO_DB_PASSWORD"),
		os.Getenv("MONGO_DB_HOST"),
		os.Getenv("MONGO_DB_PORT"),
		os.Getenv("MONGO_DB_DATABASE"),
	)

	client, err := connectMongoDatabase(dsnMongo)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	mongoClient = client

	return &MongoDBWrapper{client: mongoClient}
}

// connectMongoDatabase handles MongoDB connection setup.
func connectMongoDatabase(dsnMongo string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(dsnMongo)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	return client, nil
}

// GetDb method to satisfy the DatabaseMongo interface.
func (wrapper *MongoDBWrapper) GetDb() *mongo.Client {
	return wrapper.client
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance initializes and returns a MongoDB client instance
func DBinstance() *mongo.Client {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve MongoDB connection URL from environment variables
	MongoDb := os.Getenv("MONGODB_URL")

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	// Set a timeout context for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// Global MongoDB client instance
var Client *mongo.Client = DBinstance()

// OpenCollection returns a reference to a specified collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Connect to the database named "cluster0" and retrieve the specified collection
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}

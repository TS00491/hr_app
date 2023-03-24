package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables to store the MongoDB connection string, client, and database
var (
	mongoURI string
	client   *mongo.Client
	db       *mongo.Database
)

// init function runs when the package is imported
func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the MongoDB connection string from the environment variable
	mongoURI = os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}

	// Create a new MongoDB client
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Create a context with a 10-second timeout for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB using the client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Set the database to the "hrapp" database
	db = client.Database("hrapp")
}

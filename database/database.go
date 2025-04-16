package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Declare global variables
var DB *mongo.Database
var client *mongo.Client

func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Ping the database
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB")
	client = mongoClient
	DB = client.Database("fiberapp") // Change this to your database name
	return nil
}

// GetCollection returns a collection from the connected database
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database connection is not initialized. Call ConnectDB() first.")
	}
	return DB.Collection(collectionName)
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		} else {
			fmt.Println("Disconnected from MongoDB")
		}
	}
}

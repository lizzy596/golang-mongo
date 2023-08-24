package database

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB() (*mongo.Client, error) {
	// Replace with your MongoDB connection details
	connectionString := "mongodb://mongo:27001,mongo:27002,mongo:27003/go-server"
// Create a new MongoDB client
client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
if err != nil {
	return nil, err
}

// Create a context with a timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Connect to the MongoDB server
err = client.Connect(ctx)
if err != nil {
	return nil, err
}

// Ping the server to check the connection status
err = client.Ping(ctx, readpref.Primary())
if err != nil {
	return nil, err
}

fmt.Println("Connected to MongoDB!")

// Return the client and error (if any)
return client, nil
}



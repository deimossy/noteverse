package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Printf("failed to Connect to MongoDB: %v", err )
		return nil, err
	}

	//! Test connection
	ctxPing, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := client.Ping(ctxPing, nil); err != nil {
		log.Printf("MongoDB ping failed: %v", err)
		return nil, err
	}

	log.Println("Connected to MongoDB successfully")
	return client, nil
}
package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://mongo:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	log.Println("✅ Connected to MongoDB")
	DB = client
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("usercore").Collection(collectionName)
}

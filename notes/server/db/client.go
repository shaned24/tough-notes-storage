package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoClient() *mongo.Client {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("failed to initialize the mongodb client: %v", err)
	}

	err = mongoClient.Connect(context.TODO())
	if err != nil {
		log.Fatalf("failed to connect to mongodb server: %v", err)
	}

	return mongoClient
}

func NewCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

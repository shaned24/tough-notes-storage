package database

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoConnection struct {
	Client *mongo.Client
	Config *Config
}

func (*MongoConnection) Name() string {
	return "MongoConnection"
}

func NewMongoDB(client *mongo.Client, config *Config) *MongoConnection {
	return &MongoConnection{Client: client, Config: config}
}

func (s *MongoConnection) Connect(_ context.Context) error {
	mongoCtx, _ := context.WithTimeout(context.Background(), s.Config.ConnectionTimeout)
	if err := s.Client.Connect(mongoCtx); err != nil {
		return errors.New(fmt.Sprintf("failed to connect to mongodb server: %v", err))
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err := s.Client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.New("mongo server not available")
	}
	return nil
}

func (s *MongoConnection) Disconnect(ctx context.Context) error {
	return s.Client.Disconnect(ctx)
}

func NewMongoCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func NewMongoClient(cfg *Config) *mongo.Client {
	var mongoClient *mongo.Client
	var err error

	opts := options.Client()
	opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port))

	if mongoClient, err = mongo.NewClient(opts); err != nil {
		log.Fatalf("failed to initialize the mongodb client: %v", err)
	}

	return mongoClient
}

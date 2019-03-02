package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func NewMongoCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func NewMongoStorage(client *mongo.Client, collection *mongo.Collection) *MongoStorage {
	return &MongoStorage{
		Collection: collection,
		Client:     client,
	}
}

type noteItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type MongoStorage struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (s *MongoStorage) CreateNote(ctx context.Context, note *NoteItem) (*NoteItem, error) {
	res, err := s.Collection.InsertOne(ctx, noteItem{
		AuthorID: note.AuthorID,
		Content:  note.Content,
		Title:    note.Title,
	})

	if err != nil {
		return nil, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID: %v", err))
	}

	note.ID = oid.Hex()

	return note, nil
}

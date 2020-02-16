package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type noteItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type MongoConfig struct {
	Host              string
	Port              string
	ConnectionTimeout time.Duration
}

type MongoStorage struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (s *MongoStorage) GetNote(ctx context.Context, noteId string) (*NoteItem, error) {
	oid, err := primitive.ObjectIDFromHex(noteId)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse note id: %v", err),
		)
	}

	noteItem := &noteItem{}
	filter := bson.M{"_id": oid}

	result := s.Collection.FindOne(ctx, filter)
	if err := result.Decode(noteItem); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find note with specified id: %v", oid),
		)
	}

	return &NoteItem{
		Title:    noteItem.Title,
		AuthorID: noteItem.AuthorID,
		Content:  noteItem.Content,
		ID:       noteItem.ID.Hex(),
	}, nil
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

func NewMongoClient(cfg *MongoConfig) *mongo.Client {
	var mongoClient *mongo.Client
	var err error

	opts := options.Client()
	opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port))

	if mongoClient, err = mongo.NewClient(opts); err != nil {
		log.Fatalf("failed to initialize the mongodb client: %v", err)
	}

	mongoCtx, _ := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	if err = mongoClient.Connect(mongoCtx); err != nil {
		log.Fatalf("failed to connect to mongodb server: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Mongo server not available")

	}

	return mongoClient
}

func NewMongoCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

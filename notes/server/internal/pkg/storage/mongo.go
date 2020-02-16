package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type noteItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type MongoStorage struct {
	Collection *mongo.Collection
}

func NewMongoStorage(collection *mongo.Collection) *MongoStorage {
	return &MongoStorage{Collection: collection}
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

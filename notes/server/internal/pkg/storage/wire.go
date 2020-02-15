package storage

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var Providers = wire.NewSet(
	ProvideNoteStorage,
	ProvideMongoCollection,
	ProvideMongoDbName,
	ProvideMongoCollectionName,
	NewMongoClient,
)

type MongoDbName string
type MongoCollectionName string

func ProvideMongoDbName() MongoDbName {
	return "notes"
}

func ProvideMongoCollectionName() MongoCollectionName {
	return "notes"
}

func ProvideMongoCollection(m *mongo.Client, dbName MongoDbName, collection MongoCollectionName) *mongo.Collection {
	return NewMongoCollection(m, string(dbName), string(collection))
}

func ProvideNoteStorage(client *mongo.Client, collection *mongo.Collection) NoteStorage {
	return &MongoStorage{
		Client:     client,
		Collection: collection,
	}
}

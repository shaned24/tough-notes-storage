package notes

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

var Providers = wire.NewSet(
	ProvideNoteStorage,
	ProvideMongoCollection,
	ProvideMongoDbName,
	ProvideMongoCollectionName,
	storage.NewMongoClient,
	ProvideNoteService,
)

type MongoDbName string
type MongoCollectionName string

func ProvideNoteService(s storage.NoteStorage, c *mongo.Client) *NoteService {
	return &NoteService{mongo: c, Storage: s}
}

func ProvideMongoDbName() MongoDbName {
	return "notes"
}

func ProvideMongoCollectionName() MongoCollectionName {
	return "notes"
}

func ProvideMongoCollection(m *mongo.Client, dbName MongoDbName, collection MongoCollectionName) *mongo.Collection {
	return storage.NewMongoCollection(m, string(dbName), string(collection))
}

func ProvideNoteStorage(client *mongo.Client, collection *mongo.Collection) storage.NoteStorage {
	return &storage.MongoStorage{
		Client:     client,
		Collection: collection,
	}
}

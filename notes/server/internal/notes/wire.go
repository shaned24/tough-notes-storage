package notes

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/config"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

var Providers = wire.NewSet(
	ProvideNoteStorage,
	ProvideMongoCollection,
	ProvideMongoDbName,
	ProvideMongoCollectionName,
	ProvideMongoClient,
	ProvideNoteService,
	ProvideMongoConfig,
)

type MongoDbName string
type MongoCollectionName string

func ProvideNoteService(s storage.NoteStorage, c *mongo.Client) *NoteService {
	return &NoteService{mongo: c, Storage: s}
}

func ProvideMongoConfig() (*storage.MongoConfig, error) {
	timeout, err := strconv.Atoi(config.Getenv("MONGO_CONNECTION_TIMEOUT", "5"))
	if err != nil {
		return nil, err
	}
	return &storage.MongoConfig{
		Host:              config.Getenv("MONGO_HOST", "localhost"),
		Port:              config.Getenv("MONGO_PORT", "27017"),
		ConnectionTimeout: time.Second * time.Duration(timeout),
	}, nil
}
func ProvideMongoClient(cfg *storage.MongoConfig) *mongo.Client {
	return storage.NewMongoClient(cfg)
}

func ProvideMongoDbName() MongoDbName {
	return MongoDbName(config.Getenv("MONGO_DB_NAME", "notes"))
}

func ProvideMongoCollectionName() MongoCollectionName {
	return MongoCollectionName(config.Getenv("MONGO_COLLECTION_NAME", "notes"))
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

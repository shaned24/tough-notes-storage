package notes

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/config"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/database"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

var Providers = wire.NewSet(
	ProvideNoteStorage,
	ProvideNoteService,
	ProvideDatabaseConfig,
	ProvideDatabase,
	mongoDbProviders,
)

var mongoDbProviders = wire.NewSet(
	ProvideMongoCollection,
	ProvideMongoDbName,
	ProvideMongoCollectionName,
	ProvideMongoClient,
)

type MongoDbName string
type MongoCollectionName string

func ProvideNoteService(s storage.NoteStorage) *NoteService {
	return &NoteService{Storage: s}
}

func ProvideNoteStorage(client *mongo.Client, collection *mongo.Collection, cfg *database.Config) storage.NoteStorage {
	return storage.NewMongoStorage(client, collection, cfg)
}

func ProvideDatabase(client *mongo.Client, collection *mongo.Collection, cfg *database.Config) database.Database {
	return storage.NewMongoStorage(client, collection, cfg)
}

func ProvideDatabaseConfig() (*database.Config, error) {
	timeout, err := strconv.Atoi(config.Getenv("MONGO_CONNECTION_TIMEOUT", "5"))
	if err != nil {
		return nil, err
	}

	host := config.Getenv("MONGO_HOST", "localhost")
	port := config.Getenv("MONGO_PORT", "27017")

	return database.NewDatabaseConfig(host, port, time.Duration(timeout)), nil
}

func ProvideMongoClient(cfg *database.Config) *mongo.Client {
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

package database

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/internal/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

var Providers = wire.NewSet(
	ProvideDatabaseConnection,
	ProvideDatabaseConfig,
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

func ProvideDatabaseConnection(client *mongo.Client, cfg *Config) Connection {
	return NewMongoDB(client, cfg)
}

func ProvideDatabaseConfig() (*Config, error) {
	timeout, err := strconv.Atoi(config.Getenv("MONGO_CONNECTION_TIMEOUT", "5"))
	if err != nil {
		return nil, err
	}

	host := config.Getenv("MONGO_HOST", "localhost")
	port := config.Getenv("MONGO_PORT", "27017")

	return NewDatabaseConfig(host, port, time.Duration(timeout)), nil
}

func ProvideMongoClient(cfg *Config) *mongo.Client {
	return NewMongoClient(cfg)
}

func ProvideMongoDbName() MongoDbName {
	return MongoDbName(config.Getenv("MONGO_DB_NAME", "notes"))
}

func ProvideMongoCollectionName() MongoCollectionName {
	return MongoCollectionName(config.Getenv("MONGO_COLLECTION_NAME", "notes"))
}

func ProvideMongoCollection(m *mongo.Client, dbName MongoDbName, collection MongoCollectionName) *mongo.Collection {
	return NewMongoCollection(m, string(dbName), string(collection))
}

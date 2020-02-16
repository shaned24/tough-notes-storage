package notes

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

var Providers = wire.NewSet(
	ProvideNoteStorage,
	ProvideNoteService,
)

func ProvideNoteService(s storage.NoteStorage) *NoteService {
	return &NoteService{Storage: s}
}

func ProvideNoteStorage(collection *mongo.Collection) storage.NoteStorage {
	return storage.NewMongoStorage(collection)
}

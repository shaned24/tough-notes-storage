package notes

import (
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type NoteService struct {
	MongoClient *mongo.Client
}

var collection *mongo.Collection

type noteItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func RegisterService(server *grpc.Server, noteServer *NoteService) {
	notespb.RegisterNoteServiceServer(server, noteServer)
}

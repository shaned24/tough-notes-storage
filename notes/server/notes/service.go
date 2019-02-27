package notes

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteService struct {
	Collection *mongo.Collection
}

func (s *NoteService) CreateNote(ctx context.Context, req *notespb.CreateNoteRequest) (*notespb.CreateNoteResponse, error) {
	note := req.GetNote()

	data := noteItem{
		AuthorID: note.GetAuthorId(),
		Content:  note.GetContent(),
		Title:    note.GetTitle(),
	}

	res, err := s.Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID: %v", err))
	}

	return &notespb.CreateNoteResponse{
		Note: &notespb.Note{
			Id:       oid.Hex(),
			AuthorId: note.GetAuthorId(),
			Content:  note.GetContent(),
			Title:    note.GetTitle(),
		},
	}, nil
}

type noteItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func RegisterService(server *grpc.Server, noteServer *NoteService) {
	notespb.RegisterNoteServiceServer(server, noteServer)
}

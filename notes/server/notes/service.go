package notes

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"github.com/shaned24/tough-notes-storage/notes/server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteService struct {
	Storage storage.NoteStorage
}

func (s *NoteService) ReadNote(context.Context, *notespb.ReadNoteRequest) (*notespb.ReadNoteResponse, error) {
	panic("implement me")
}

func (s *NoteService) CreateNote(ctx context.Context, req *notespb.CreateNoteRequest) (*notespb.CreateNoteResponse, error) {
	note := req.GetNote()

	noteItem := &storage.NoteItem{
		AuthorID: note.GetAuthorId(),
		Content:  note.GetContent(),
		Title:    note.GetTitle(),
	}

	noteItem, err := s.Storage.CreateNote(ctx, noteItem)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err))
	}

	return &notespb.CreateNoteResponse{
		Note: &notespb.Note{
			Id:       noteItem.ID,
			AuthorId: noteItem.AuthorID,
			Content:  noteItem.Content,
			Title:    noteItem.Title,
		},
	}, nil
}

func RegisterService(server *grpc.Server, noteServer *NoteService) {
	notespb.RegisterNoteServiceServer(server, noteServer)
}

func NewNoteService(s storage.NoteStorage) *NoteService {
	return &NoteService{
		Storage: s,
	}
}

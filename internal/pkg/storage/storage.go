package storage

import "context"

type NoteStorage interface {
	CreateNote(context.Context, *NoteItem) (*NoteItem, error)
	GetNote(context.Context, string) (*NoteItem, error)
}

type NoteItem struct {
	ID       string
	AuthorID string
	Content  string
	Title    string
}

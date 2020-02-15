//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/notes"
)

func setupNoteService() (*notes.NoteService, error) {

	wire.Build(
		notes.Providers,
	)
	return nil, nil
}

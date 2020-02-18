package main

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/internal/notes"
	"github.com/shaned24/tough-notes-storage/internal/pkg/database"
	"github.com/shaned24/tough-notes-storage/internal/server"
	"log"
)

var wireProviders = wire.NewSet(
	server.Providers,
	database.Providers,
	notes.Providers,
)

func main() {
	s, _ := setupServer()

	err := s.Start()
	if err != nil {
		log.Fatalf("Encountered an error while starting the server: %v", err)
	}
}

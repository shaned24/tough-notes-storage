package main

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/internal/pkg/database"
	"github.com/shaned24/tough-notes-storage/internal/pkg/notes"
	"github.com/shaned24/tough-notes-storage/internal/server"
	"log"
	"os"
	"os/signal"
)

var wireProviders = wire.NewSet(
	server.Providers,
	database.Providers,
	notes.Providers,
)

func main() {
	s, _ := setupServer()

	go func() {
		err := s.Start()
		if err != nil {
			log.Fatalf("Encountered an error while starting the server: %v", err)
		}
	}()

	go func() {
		err := s.HttpGateway.Start()
		if err != nil {
			log.Fatalf("Encountered an error while starting the server: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)
	<-wait

	log.Println("Shutting down main.")

}

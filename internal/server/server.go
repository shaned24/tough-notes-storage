package server

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/api/notespb"
	"github.com/shaned24/tough-notes-storage/internal/notes"
	"github.com/shaned24/tough-notes-storage/internal/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type Config struct {
	Host     string
	Port     string
	Protocol string
}

type Server struct {
	Config       Config
	NotesService *notes.NoteService
	Database     database.Database
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.Config.Protocol, fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port))
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	gRPCServer := grpc.NewServer(opts...)

	if err = s.connectToDatabase(); err != nil {
		return err
	}

	// register the gRPC server with our protocol buffers
	notespb.RegisterNoteServiceServer(gRPCServer, s.NotesService)

	go func() {
		log.Println("Serving on", fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port))
		log.Printf("tough notes service started.")
		err = gRPCServer.Serve(listener)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)
	<-wait

	log.Println("Shutting down server.")
	gRPCServer.Stop()
	log.Println("Stopping the listener.")
	_ = listener.Close()
	log.Println("Stopping Notes service.")

	if err = s.disconnectFromDatabase(); err != nil {
		return err
	}

	return nil
}

// Connect to the Database if one exists
func (s *Server) connectToDatabase() error {

	if s.Database != nil {
		log.Printf("Trying to connect to %s...", s.Database.Name())
		err := s.Database.Connect(context.Background())
		if err != nil {
			return err
		}
		log.Printf("Connected to %s.", s.Database.Name())

	}
	return nil
}

// disconnect to the Database if one exists
func (s *Server) disconnectFromDatabase() error {
	if s.Database != nil {
		log.Printf("Disconnecting from %s...", s.Database.Name())
		err := s.Database.Disconnect(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

package server

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/api/notespb"
	"github.com/shaned24/tough-notes-storage/internal/pkg/database"
	"github.com/shaned24/tough-notes-storage/internal/pkg/gateway"
	"github.com/shaned24/tough-notes-storage/internal/pkg/notes"
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
	Config             Config
	NotesService       *notes.NoteService
	DatabaseConnection database.Connection
	HttpGateway        *gateway.HttpGateway
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.Config.Protocol, fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port))
	if err != nil {
		return err
	}

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(LogRequestInfoMiddleware()),
	}

	gRPCServer := grpc.NewServer(serverOptions...)

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

// Connect to the Connection if one exists
func (s *Server) connectToDatabase() error {

	if s.DatabaseConnection != nil {
		log.Printf("Trying to connect to %s...", s.DatabaseConnection.Name())
		err := s.DatabaseConnection.Connect(context.Background())
		if err != nil {
			return err
		}
		log.Printf("Connected to %s.", s.DatabaseConnection.Name())

	}
	return nil
}

// disconnect to the Connection if one exists
func (s *Server) disconnectFromDatabase() error {
	if s.DatabaseConnection != nil {
		log.Printf("Disconnecting from %s...", s.DatabaseConnection.Name())
		err := s.DatabaseConnection.Disconnect(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

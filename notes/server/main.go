package main

import (
	"fmt"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	// if we crash the go code, the log shows the file name and line number of an error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	protocol := "tcp"
	host := "0.0.0.0"
	port := "50051"

	listener, err := net.Listen(protocol, fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	service, _ := setupNoteService()
	notespb.RegisterNoteServiceServer(s, service)

	go func() {
		log.Println("Serving on", fmt.Sprintf("%s:%s", host, port))
		log.Printf("tough notes service started.")
		err = s.Serve(listener)

		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)
	<-wait

	log.Println("Shutting down server.")
	s.Stop()
	log.Println("Stopping the listener.")
	_ = listener.Close()
	log.Println("Closing mongodb connection.")
	_ = service.Stop()
}
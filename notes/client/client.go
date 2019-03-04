package main

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	host := "localhost"
	port := "50051"

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := notespb.NewNoteServiceClient(conn)

	doCreateNoteUnary(client)
	doReadNoteUnary(client)
}

func doCreateNoteUnary(client notespb.NoteServiceClient) {
	log.Println("Creating a note...")

	req := &notespb.CreateNoteRequest{
		Note: &notespb.Note{
			Content:  "# Some Markdown yes!",
			AuthorId: "shane_daly",
			Title:    "First Note",
		},
	}

	resp, err := client.CreateNote(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", resp.Note)
}

func doReadNoteUnary(client notespb.NoteServiceClient) {
	log.Println("Reading a note...")

	req := &notespb.ReadNoteRequest{
		Id: "5c7c7bc6bfc49da93acfefa9",
	}

	resp, err := client.ReadNote(context.Background(), req)

	if err != nil {
		statusErr, ok := status.FromError(err)

		if ok {
			switch statusErr.Code() {
			case codes.InvalidArgument:
				log.Fatalf("Invalid argument passed: %v", err)
			case codes.NotFound:
				log.Fatalf("Note not found: %v", err)
			default:
				log.Fatalf("Unexpected error: %v", err)
			}
		} else {
			log.Fatalf("error while calling ReadNote RPC: %v", err)
		}
	}

	log.Printf("Response from ReadNote Request: %v", resp.Note)
}

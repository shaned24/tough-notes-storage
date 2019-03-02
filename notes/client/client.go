package main

import (
	"context"
	"fmt"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"google.golang.org/grpc"
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

	doUnary(client)
}

func doUnary(client notespb.NoteServiceClient) {
	log.Println("doing unary request")

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

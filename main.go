package main

import "log"

func main() {
	server, _ := setupServer()

	err := server.Start()
	if err != nil {
		log.Fatalf("Encountered an error while starting the server: %v", err)
	}
}

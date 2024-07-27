package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Use FileServer to serve static files
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/", fileServer)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

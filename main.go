package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Start the server
	server.ListenAndServe()
}

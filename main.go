package main

import (
	"log"
	"net/http"
	"time"
)

// Response structure for JSON responses
type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// main initializes and starts the server
func main() {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/", handlers.NotFoundHandler)

	// Start server
	server := &http.Server{
		Addr:         ":8069",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8069")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

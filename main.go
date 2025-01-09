package main

import (
	"log"
	"net/http"
  "math/rand"
	"time"
)

// Response structure for JSON responses
type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

urls := make(map[string]string)

// GenerateRandomString generates a random string of the given length.
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func ShortenHandler(url: string) {
  urls[string] := GenerateRandomString(8)
  return urls[string]
}

// main initializes and starts the server
func main() {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/", handlers.NotFoundHandler)
  mux.HandleFunc("/shorten/", ShortenHandler)

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

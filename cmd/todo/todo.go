package main

import (
	"github.com/arunkpatra/todo/internal/app/todo/handlers"
	"net/http"
)

func main() {
	// Handle requests
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)

	// Start the server
	_ = http.ListenAndServe(":8080", nil)
}
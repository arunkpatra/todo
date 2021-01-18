package main

import (
	"fmt"
	"github.com/arunkpatra/todo/internal/app/todo/db"
	"github.com/arunkpatra/todo/internal/app/todo/handlers"
	"net/http"
)

func main() {
	// Setup database
	_, err := db.SetupDatabase()

	if err != nil {
		fmt.Println("Database setup failed.")
	}

	// Handle requests
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)

	// Start the server
	_ = http.ListenAndServe(":8080", nil)
}

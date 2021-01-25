package main

import (
	"fmt"
	"github.com/arunkpatra/todo/internal/app/todo/db"
	"github.com/arunkpatra/todo/internal/app/todo/handlers"
	"github.com/arunkpatra/todo/internal/app/todo/service"
	"net/http"
)

func main() {
	// Setup database
	thisDB, err := db.SetupDatabase()

	defer thisDB.Close()

	// Setup the service
	service.MyTodoService = &service.TodoService{DB: thisDB}

	if err != nil {
		fmt.Println("Database setup failed.")
	}

	// Handle requests
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/todo", handlers.TodoHandler)

	// Start the server
	_ = http.ListenAndServe(":8080", nil)
}

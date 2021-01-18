package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// SetupDatabase sets up the access to the database
func SetupDatabase() (*sql.DB, error) {
	// Get a handle to the database. sql.Open() doesn't directly open a connection.
	db, err := sql.Open("mysql", "root:Password123456@tcp(127.0.0.1:3306)/todoDB?checkConnLiveness=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Setup database successfully.")
	return db, err
}

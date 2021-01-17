package main

import (
	"fmt"
	"net/http"
)

// hello writes a response for the /hello path
func hello(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there!\n")
}

func main() {
	// Handle /hello
	http.HandleFunc("/hello", hello)

	// Start the server
	_ = http.ListenAndServe(":8080", nil)
}

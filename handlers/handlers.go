package handlers

import (
	"io"
	"net/http"
)

// HealthHandler handles the /health route. It returns an application status in JSOn format.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, _ = io.WriteString(w, `{"status" : "UP"}`)
}

// HelloHandler returns a friendly message. This is just a demo method.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = io.WriteString(w, "Hi there!")
}

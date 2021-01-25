package handlers

import (
	"encoding/json"
	"github.com/arunkpatra/todo/internal/app/todo/service"
	"io"
	"net/http"
)

// TodoHandler handles the /todo route.
func TodoHandler(w http.ResponseWriter, r *http.Request) {

	td, _ := service.MyTodoService.Read(r.Context())
	x, _ := json.Marshal(td)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, _ = io.WriteString(w, string(x))

}

// HealthHandler handles the /health route. It returns an application status in JSOn format.
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, _ = io.WriteString(w, `{"status" : "UP"}`)
}

// HelloHandler returns a friendly message. This is just a demo method.
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	hr, _ := service.MyTodoService.Hello(r.Context())
	x, _ := json.Marshal(hr)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = io.WriteString(w, string(x))
}

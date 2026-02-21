package main

import (
	"go-rest-api-basics/cmd/handlers"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HeeloRequest struct {
	Name string `json:"name"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/hello", HelloHandler)
	r.Post("/hello", HelloPostHandler)
	r.Get("/todos", handlers.GetTodosHandler)
	r.Post("/todos", handlers.CreateTodoHandler)
	r.Put("/todos/{id}",handlers.UpdateTodoHandler)
	r.Delete("/todos/{id}",handlers.DeleteTodoHandler)
	http.ListenAndServe(":8080", r)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, World!",
	})
}

func HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req HeeloRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, " + req.Name + "!",
	})
}

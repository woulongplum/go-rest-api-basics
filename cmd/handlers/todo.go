package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-rest-api-basics/cmd/models"

	"github.com/go-chi/chi/v5"
)

var todos []models.Todo
var nextID = 1

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	todo := models.Todo{ID: nextID, Text: req.Text, Done: false}
	nextID++
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	todoID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var targetIndex = -1
	for i, t := range todos {
		if t.ID == todoID {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	var req struct {
		Text string `json:"text"`
		Done bool `json:"done"`
	}
	

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	todos[targetIndex].Text = req.Text
	todos[targetIndex].Done = req.Done

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos[targetIndex])
}


func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idPram := chi.URLParam(r,"id")

	todoID , err := strconv.Atoi(idPram)
	if err != nil {
		http.Error(w,"Invalid ID", http.StatusBadRequest)
		return
	}

	targetIndex := -1
	for i, t := range todos {
		if t.ID == todoID {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		http.Error(w,"Todo not found", http.StatusNotFound)
		return
	}

	todos = append(todos[:targetIndex], todos[targetIndex+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-rest-api-basics/cmd/models"

	"github.com/go-chi/chi/v5"
)

var todos []models.Todo
var nextID = 1

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var req struct {
		Text string `json:"text"`	
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//テキストが空白のみの場合はエラー　バリデーションチェック
	if strings.TrimSpace(req.Text) == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	//DBにINSERT
	result , err := DB.Exec(
		"INSERT INTO todos (text, done) VALUES (?,?)",
		req.Text,
		false,
	)

	if err != nil {
		http.Error(w,"DB error", http.StatusInternalServerError)
		return
	}
	
	id , _ := result.LastInsertId()

	todo := models.Todo{
		ID: int(id),
		Text: req.Text,
		Done: false,
	}

	json.NewEncoder(w).Encode(todo)
} 

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	rows, err := DB.Query("SELECT id, text, done FROM todos")
	if err != nil {
		http.Error(w,"DB query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID,&t.Text,&t.Done); err != nil {
			http.Error(w,"DB scan error", http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		 http.Error(w, "DB rows error", http.StatusInternalServerError) 
		 return 
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil { 
		http.Error(w, "JSON encode error", http.StatusInternalServerError) 
		return 
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")



	idParam := chi.URLParam(r, "id")

	todoID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
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

	if strings.TrimSpace(req.Text) == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	result , err := DB.Exec(
		"UPDATE todos SET text = ?, done = ? WHERE id = ? ",
		req.Text,
		req.Done,
		todoID,
	)

	if err != nil {
		http.Error(w, "DB update error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return	
	}

	updated := models.Todo {
		ID: todoID,
		Text: req.Text,
		Done: req.Done,	
	}

	json.NewEncoder(w).Encode(updated)

	
}


func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idPram := chi.URLParam(r,"id")

	todoID , err := strconv.Atoi(idPram)
	if err != nil {
		http.Error(w,"Invalid ID", http.StatusBadRequest)
		return
	}

	result , err := DB.Exec("DELETE FROM todos WHERE id = ?",todoID)
	if err != nil {
		http.Error(w,"DB delete error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return	
	}
	
	w.WriteHeader(http.StatusNoContent)
}

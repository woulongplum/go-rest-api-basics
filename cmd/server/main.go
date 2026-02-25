package main

import (
	"go-rest-api-basics/cmd/handlers"
	"go-rest-api-basics/internal/db"
	"log"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HeeloRequest struct {
	Name string `json:"name"`
}

func main() {

	//DB 接続
	conn,err := db.ConnectPostgreSQL()
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}

	//テーブル作成（初期化
	if err := db.InitPostgreSQL(conn); err != nil {
		log.Fatal("テーブル作成エラー:", err)
	}

	//DBをハンドラーにセット
	handlers.SetDB(conn)

	//ルーター設定
	r := chi.NewRouter()

	r.Get("/hello", HelloHandler)
	r.Post("/hello", HelloPostHandler)
	r.Get("/todos", handlers.GetTodosHandler)
	r.Post("/todos", handlers.CreateTodoHandler)
	r.Put("/todos/{id}",handlers.UpdateTodoHandler)
	r.Delete("/todos/{id}",handlers.DeleteTodoHandler)

	//サーバー起動
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

package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//PostgreSQL に接続する
func ConnectPostgreSQL() (*sql.DB, error) {
	return  sql.Open("postgres","host=localhost port=5432 user=postgres password=postgres dbname=todoapp sslmode=disable")
}

//初期化（テーブル作成）
func InitPostgreSQL(db *sql.DB) error { 
	query := ` CREATE TABLE IF NOT EXISTS todos ( id SERIAL PRIMARY KEY, text TEXT NOT NULL, done BOOLEAN NOT NULL DEFAULT false ); `
	 _, err := db.Exec(query) 
	 return err 
}

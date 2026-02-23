package db

import (
	"database/sql"
	

	_ "github.com/mattn/go-sqlite3"
)

//SQLite に接続する
func ConnectSQLite() (*sql.DB, error) {
	return sql.Open("sqlite3", "./internal/data/todos.db")
}

//初期化（テーブル作成）
func InitSQLite(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT NOT NULL,
	done BOOLEAN NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

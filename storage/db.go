package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPATH = "./database/todo.sqlite"
)

// Init initializes the database
func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBPATH)
	if err != nil {
		return db, err
	}

	createStatement := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NULL,
			content TEXT NOT NULL,
			created_at DATETIME NULL,
			completed_at DATEIME NULL
		)
	`
	if _, err = db.Exec(createStatement); err != nil {
		return db, err
	}

	return db, nil
}

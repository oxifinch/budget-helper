package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Database wraps an SQL driver with an open connection that can
// be passed around to other functions.
type Database struct {
	*sql.DB
}

func NewDatabase() *Database {
	conn, err := sql.Open("sqlite3", "./app-db.db")
	if err != nil {
		log.Fatalf("NewDatabase: failed to connect to database: %v", err)
	}

	return &Database{conn}
}

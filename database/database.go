package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// A Database is a struct that (currently) contains one member variable:
// a database connection, which can be passed around and used by other repos.
type Database struct {
	*sql.DB
}

func NewDatabase() *Database {
	conn, err := sql.Open("sqlite3", "./app-db.db")
	if err != nil {
		log.Fatalf("error: failed to connect to database: %v", err)
	}

	return &Database{conn}
}

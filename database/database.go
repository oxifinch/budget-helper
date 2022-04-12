package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// A Database is a struct that (currently) contains ONE member variable:
// a database connection, which can be passed around and used by other repos.
// TODO: This Database should be a wrapper around SQLBoiler, not just a regular
// SQL database connection thingy.
type Database struct {
	*sql.DB
}

// TODO: Should this return a pointer to a Database, or just a Database?
func NewDatabase() *Database {
	conn, err := sql.Open("sqlite3", "./app-db.db")
	if err != nil {
		log.Printf("ERROR: Failed to connect to database. Aborting...\n")
		os.Exit(1)
	}

	return &Database{conn}
}

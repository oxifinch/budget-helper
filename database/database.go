package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database wraps an SQL driver with an open connection that can
// be passed around to other functions.
type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("./app-db.db"), &gorm.Config{})
	// conn, err := sql.Open("sqlite3", "./app-db.db")
	if err != nil {
		log.Fatalf("NewDatabase: failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&User{}, &Budget{})
	if err != nil {
		log.Fatalf("NewDatabase: %v\n", err)
	}

	return &Database{db}
}

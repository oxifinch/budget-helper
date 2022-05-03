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

func (db *Database) Seed() {
	// Delete old data first so there are no duplicates.
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM budgets")

	newUsers := []*User{
		{Username: "joseph", Password: "secret01"},
		{Username: "jean-paul", Password: "secret02"},
		{Username: "bubby", Password: "secret03"},
	}
	for _, u := range newUsers {
		err := db.Create(u).Error
		if err != nil {
			log.Fatalf("seed: %v\n", err)
		}
	}

	newBudgets := []*Budget{
		{StartDate: "2022-04-25", EndDate: "2022-05-25", UserID: 1},
		{StartDate: "2022-03-28", EndDate: "2022-04-24", UserID: 2},
		{StartDate: "2022-05-10", EndDate: "2022-06-10", UserID: 3},
	}
	for _, b := range newBudgets {
		err := db.Create(b).Error
		if err != nil {
			log.Fatalf("seed: %v\n", err)
		}
	}
}

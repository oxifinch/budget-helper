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

	err = db.AutoMigrate(&User{}, &Budget{}, &Category{})
	if err != nil {
		log.Fatalf("NewDatabase: %v\n", err)
	}

	return &Database{db}
}

func (db *Database) Seed() {
	// Delete old data first so there are no duplicates.
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM budgets")
	db.Exec("DELETE FROM categories")

	newUsers := []*User{
		{Username: "joseph", Password: "secret01"},
		{Username: "jean-paul", Password: "secret02"},
		{Username: "bubby", Password: "secret03"},
	}
	for _, u := range newUsers {
		err := db.Create(u).Error
		if err != nil {
			log.Fatalf("Seed: %v\n", err)
		}
	}

	newBudgets := []*Budget{
		{StartDate: "2022-04-25", EndDate: "2022-05-25", Allocated: 11500.00, Currency: "SEK", UserID: 1},
		{StartDate: "2022-03-28", EndDate: "2022-04-24", Allocated: 9250.99, Currency: "SEK", UserID: 2},
		{StartDate: "2022-05-10", EndDate: "2022-06-10", Allocated: 15000.00, Currency: "SEK", UserID: 3},
	}
	for _, b := range newBudgets {
		err := db.Create(b).Error
		if err != nil {
			log.Fatalf("Seed: %v\n", err)
		}
	}

	newCategories := []*Category{
		{Name: "Groceries", Description: "Food and stuff.", Color: "#b6eea6", UserID: 1},
		{Name: "Entertainment", Description: "Videogames and whatever.", Color: "#a4778b", UserID: 1},
		{Name: "Learning Material", Description: "Books and courses!", Color: "#a9ffcb", UserID: 1},
	}
	for _, c := range newCategories {
		err := db.Create(c).Error
		if err != nil {
			log.Fatalf("Seed: %v\n", err)
		}
	}
}

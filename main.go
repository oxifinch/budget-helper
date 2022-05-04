package main

import (
	"budget-helper/database"
)

const (
	debug = true
	port  = ":8080"
)

func main() {
	db := database.NewDatabase()
	if debug {
		db.Seed()
	}

	server := NewServer(port, db)
	server.Run()
}

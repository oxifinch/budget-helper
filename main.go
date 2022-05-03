package main

import (
	"budget-helper/database"
)

const port = ":8080"

func main() {
	db := database.NewDatabase()
	// TODO: GORM doesn't appear to have any close method(something about connection pooling.) Should I simply remove this?
	// defer db.Close()

	server := NewServer(port, db)
	server.Run()
}

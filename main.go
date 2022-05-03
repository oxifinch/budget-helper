package main

import (
	"budget-helper/database"
)

const port = ":8080"

func main() {
	db := database.NewDatabase()

	server := NewServer(port, db)
	server.Run()

}

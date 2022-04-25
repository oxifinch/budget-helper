package main

import (
	"budget-helper/database"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	db := database.NewDatabase()
	defer db.Close()

	router := initRouter(db)

	log.Printf("Server started on PORT %v...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

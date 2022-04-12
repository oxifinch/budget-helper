package main

import (
	"budget-helper/database"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	router := initRouter()

	db := database.NewDatabase()
	defer db.Close()

	log.Printf("Server started on PORT %v...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

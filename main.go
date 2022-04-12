package main

import (
	"budget-helper/database"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	// Initialization steps:
	// - [x] Register routes in the http.ServeMux called "router"
	// - [x] Connect to database
	// - [x] Listen for connections on port
	router := initRouter()

	db := database.NewDatabase()
	defer db.Close()

	log.Printf("Server started on PORT %v...\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("error: failed to start listener: %v", err)
	}
}

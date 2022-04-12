package main

import (
	"budget-helper/database"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	// Initialization steps:
	// - [x] Register routes in the http.ServeMux called "router"
	// - [x] Connect to database
	// - [x] Listen for connections on PORT
	router := initRouter()

	db := database.NewDatabase()
	defer db.Close()

	log.Printf("Server started on PORT %v...\n", PORT)
	http.ListenAndServe(PORT, router)
}

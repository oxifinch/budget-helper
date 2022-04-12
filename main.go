package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	// Initialization steps:
	// - [x] Register routes in the http.ServeMux called "router"
	// - [ ] Connect to database
	// - [x] Listen for connections on PORT
	router := initRouter()

	fmt.Printf("Server started on PORT %v...\n", PORT)
	http.ListenAndServe(PORT, router)
}

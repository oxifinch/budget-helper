package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	router := initRouter()

	log.Printf("Server started on PORT %v...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

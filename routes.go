package main

import (
	"fmt"
	"log"
	"net/http"
)

// Routes are named as "routeDOMAIN"

func routeHome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Panic(err)
	}
}

func initRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", routeHome)

	return router
}

package main

import (
	"budget-helper/database"
	"fmt"
	"log"
	"net/http"
)

func home(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Panic(err)
	}
}

func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func initRouter() *http.ServeMux {
	// INITIALIZE DATABASE AND ROUTES
	db := database.NewDatabase()
	router := http.NewServeMux()

	// REGISTER ALL ROUTES HERE
	router.HandleFunc("/", wrapWithDB(db, home))

	defer db.Close()
	return router
}

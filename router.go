package main

import (
	"budget-helper/database"
	"budget-helper/routes"
	"net/http"
)

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
	router.HandleFunc("/", wrapWithDB(db, routes.Home))

	defer db.Close()
	return router
}

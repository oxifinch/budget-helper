package router

import (
	"budget-helper/database"
	"log"
	"net/http"
)

type Router struct {
	Handler *http.ServeMux
}

func NewRouter(db *database.Database) *Router {
	err := db.Ping()
	if err != nil {
		log.Fatalf("NewRouter: %v\n", err)
	}

	h := http.NewServeMux()

	// -- REGISTER ALL ROUTES HERE --
	h.HandleFunc("/", wrapWithDB(db, handleHome))
	h.HandleFunc("/users", wrapWithDB(db, handleUsers))

	return &Router{Handler: h}
}

func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

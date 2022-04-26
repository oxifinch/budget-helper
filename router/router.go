package router

import (
	"budget-helper/database"
	"budget-helper/users"
	"net/http"
)

type Router struct {
	Handler *http.ServeMux
}

func NewRouter(u *users.UserRepo) *Router {
	h := http.NewServeMux()

	// -- REGISTER ALL ROUTES HERE --
	h.HandleFunc("/", handleHome)
	h.HandleFunc("/users", wrapWithDB(u.DB, handleUsers))

	return &Router{Handler: h}
}

func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

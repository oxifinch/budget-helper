package router

import (
	"budget-helper/users"
	"net/http"
)

type Router struct {
	Handler  *http.ServeMux
	UserRepo *users.UserRepo
}

func NewRouter(u *users.UserRepo) *Router {
	h := http.NewServeMux()

	return &Router{
		Handler:  h,
		UserRepo: u,
	}
}

func (rt *Router) RegisterRoutes() {
	// -- REGISTER ALL ROUTES HERE --
	rt.Handler.HandleFunc("/", rt.handleHome)
	rt.Handler.HandleFunc("/users", rt.handleUsers)
}

//func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		f(db, w, r)
//	}
//}

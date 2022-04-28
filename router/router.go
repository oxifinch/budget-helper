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
	rt.Handler.HandleFunc("/login", rt.handleLogin)
	rt.Handler.HandleFunc("/loginSave", rt.handleLoginSave)
	rt.Handler.HandleFunc("/register", rt.handleRegister)
	rt.Handler.HandleFunc("/registerSave", rt.handleRegisterSave)
	rt.Handler.HandleFunc("/dashboard", rt.handleDashboard)
}

package router

import (
	"budget-helper/budgets"
	"budget-helper/categories"
	"budget-helper/users"
	"net/http"
)

type Router struct {
	Handler      *http.ServeMux
	UserRepo     *users.UserRepo
	BudgetRepo   *budgets.BudgetRepo
	CategoryRepo *categories.CategoryRepo
}

func NewRouter(u *users.UserRepo, b *budgets.BudgetRepo, c *categories.CategoryRepo) *Router {
	h := http.NewServeMux()

	return &Router{
		Handler:      h,
		UserRepo:     u,
		BudgetRepo:   b,
		CategoryRepo: c,
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

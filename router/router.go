package router

import (
	"budget-helper/budgets"
	"budget-helper/categories"
	"budget-helper/payments"
	"budget-helper/users"
	"net/http"
)

type Router struct {
	Handler      *http.ServeMux
	UserRepo     *users.UserRepo
	BudgetRepo   *budgets.BudgetRepo
	CategoryRepo *categories.CategoryRepo
	PaymentRepo  *payments.PaymentRepo
}

func NewRouter(u *users.UserRepo, b *budgets.BudgetRepo,
	c *categories.CategoryRepo, p *payments.PaymentRepo) *Router {
	h := http.NewServeMux()

	return &Router{
		Handler:      h,
		UserRepo:     u,
		BudgetRepo:   b,
		CategoryRepo: c,
		PaymentRepo:  p,
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
	rt.Handler.HandleFunc("/newBudget", rt.handleNewBudget)
	rt.Handler.HandleFunc("/newBudgetSave", rt.handleNewBudgetSave)
	rt.Handler.HandleFunc("/payments/budget", rt.handlePaymentsBudget)
	rt.Handler.HandleFunc("/payments/budgetcategory", rt.handlePaymentsBudgetCategory)
	rt.Handler.HandleFunc("/payments/user", rt.handlePaymentsUser)
	rt.Handler.HandleFunc("/paymentSave", rt.handlePaymentSave)
}

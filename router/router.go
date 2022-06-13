package router

import (
	"budget-helper/budgets"
	"budget-helper/categories"
	"budget-helper/incomeexpenses"
	"budget-helper/payments"
	"budget-helper/users"
	"net/http"

	"github.com/gorilla/sessions"
)

type Router struct {
	Handler           *http.ServeMux
	UserRepo          *users.UserRepo
	BudgetRepo        *budgets.BudgetRepo
	CategoryRepo      *categories.CategoryRepo
	PaymentRepo       *payments.PaymentRepo
	IncomeExpenseRepo *incomeexpenses.IncomeExpenseRepo
	Store             *sessions.CookieStore
}

func NewRouter(u *users.UserRepo, b *budgets.BudgetRepo,
	c *categories.CategoryRepo, p *payments.PaymentRepo,
	i *incomeexpenses.IncomeExpenseRepo, st *sessions.CookieStore) *Router {
	h := http.NewServeMux()

	return &Router{
		Handler:           h,
		UserRepo:          u,
		BudgetRepo:        b,
		CategoryRepo:      c,
		PaymentRepo:       p,
		IncomeExpenseRepo: i,
		Store:             st,
	}
}

func (rt *Router) RegisterRoutes() {
	// == REGISTER ALL ROUTES HERE ==
	rt.Handler.HandleFunc("/", rt.handleHome)
	rt.Handler.HandleFunc("/dashboard", rt.handleDashboard)

	// - Authentication -
	rt.Handler.HandleFunc("/login", rt.handleLogin)
	rt.Handler.HandleFunc("/loginSave", rt.handleLoginSave)
	rt.Handler.HandleFunc("/register", rt.handleRegister)
	rt.Handler.HandleFunc("/registerSave", rt.handleRegisterSave)

	// - User settings -
	rt.Handler.HandleFunc("/settings", rt.handleSettings)
	rt.Handler.HandleFunc("/settings/account", rt.handleSettingsAccount)
	rt.Handler.HandleFunc("/settings/incomeexpenses", rt.handleSettingsIncomeExpenses)
	rt.Handler.HandleFunc("/settings/data/incomeexpenses", rt.handleSettingsDataIncomeExpenses)
	rt.Handler.HandleFunc("/settingsSave/account", rt.handleSettingsSaveAccount)

	// - Budget management -
	rt.Handler.HandleFunc("/newBudget", rt.handleNewBudget)
	rt.Handler.HandleFunc("/newBudgetSave", rt.handleNewBudgetSave)

	// - Payment management -
	rt.Handler.HandleFunc("/payments/budget", rt.handlePaymentsBudget)
	rt.Handler.HandleFunc("/payments/budgetcategory", rt.handlePaymentsBudgetCategory)
	rt.Handler.HandleFunc("/payments/category", rt.handlePaymentsCategory)
	rt.Handler.HandleFunc("/payments/user", rt.handlePaymentsUser)
	rt.Handler.HandleFunc("/paymentSave", rt.handlePaymentSave)

	// - Income and expenses management -
	rt.Handler.HandleFunc("/incomeexpensesCreate", rt.handleIncomeExpensesCreate)
	rt.Handler.HandleFunc("/incomeexpensesUpdate", rt.handleIncomeExpensesUpdate)
	rt.Handler.HandleFunc("/incomeexpensesDelete", rt.handleIncomeExpensesDelete)
}

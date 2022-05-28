package main

import (
	"budget-helper/budgets"
	"budget-helper/categories"
	"budget-helper/database"
	"budget-helper/incomeexpenses"
	"budget-helper/payments"
	"budget-helper/router"
	"budget-helper/users"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type Server struct {
	Router *router.Router
	Port   string
	Store  *sessions.CookieStore
}

func NewServer(port string, db *database.Database) *Server {
	sessionkey, set := os.LookupEnv("BUDGET_HELPER_SESSION_KEY")
	if !set {
		log.Fatalf("NewServer: BUDGET_HELPER_SESSION_KEY not set.\n")
	}

	st := sessions.NewCookieStore([]byte(sessionkey))

	u := users.NewUserRepo(db)
	b := budgets.NewBudgetRepo(db)
	c := categories.NewCategoryRepo(db)
	p := payments.NewPaymentRepo(db)
	i := incomeexpenses.NewIncomeExpenseRepo(db)

	router := router.NewRouter(u, b, c, p, i, st)
	router.RegisterRoutes()

	return &Server{
		Router: router,
		Port:   port,
	}
}

func (s *Server) Run() {
	log.Printf("Server starting on PORT %v\n", s.Port[1:])
	log.Fatal(http.ListenAndServe(s.Port, s.Router.Handler))
}

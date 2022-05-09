package main

import (
	"budget-helper/budgets"
	"budget-helper/categories"
	"budget-helper/database"
	"budget-helper/payments"
	"budget-helper/router"
	"budget-helper/users"
	"log"
	"net/http"
)

type Server struct {
	Router *router.Router
	Port   string
}

func NewServer(port string, db *database.Database) *Server {
	u := users.NewUserRepo(db)
	b := budgets.NewBudgetRepo(db)
	c := categories.NewCategoryRepo(db)
	p := payments.NewPaymentRepo(db)

	router := router.NewRouter(u, b, c, p)
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

package main

import (
	"budget-helper/budgets"
	"budget-helper/database"
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
	userRepo := users.NewUserRepo(db)
	budgetRepo := budgets.NewBudgetRepo(db)

	router := router.NewRouter(userRepo, budgetRepo)
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

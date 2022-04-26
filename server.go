package main

import (
	"budget-helper/database"
	"budget-helper/router"
	"budget-helper/users"
	"log"
	"net/http"
)

type Server struct {
	UserRepo *users.UserRepo
	Router   *router.Router
	Port     string
}

func NewServer(port string, db *database.Database) *Server {
	userRepo := users.NewUserRepo(db)
	router := router.NewRouter(userRepo)

	// TODO: Check that port is a valid numerical string
	return &Server{
		Router:   router,
		UserRepo: userRepo,
		Port:     port,
	}

}

func (s *Server) Run() {
	log.Printf("Server starting on PORT %v\n", s.Port[1:])
	log.Fatal(http.ListenAndServe(s.Port, s.Router.Handler))
}

package main

import (
	"budget-helper/database"
	"budget-helper/router"
	"budget-helper/users"
	"log"
	"net/http"
)

type Server struct {
	Database *database.Database
	Router   *router.Router
	UserRepo *users.UserRepo
	Port     string
}

func NewServer(port string) *Server {
	db := database.NewDatabase()
	router := router.NewRouter(db)
	userRepo := users.NewUserRepo(db)

	// TODO: Check that port is a valid numerical string
	return &Server{
		Database: db,
		Router:   router,
		UserRepo: userRepo,
		Port:     port,
	}

}

func (s *Server) Run() {
	log.Printf("Server starting on PORT %v\n", s.Port[1:])
	log.Fatal(http.ListenAndServe(s.Port, s.Router.Handler))
}

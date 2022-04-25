package main

import (
	"budget-helper/database"
	"budget-helper/users"
	"fmt"
	"log"
	"net/http"
)

func handleHome(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}
}

func handleUsers(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Fetching users...\n")
	if err != nil {
		log.Fatalf("handleUsers: %v\n", err)
	}

	repo := users.NewUserRepo(db)
	userList, err := repo.GetAllUsers()
	if err != nil {
		log.Fatalf("GetAllUsers: %v\n", err)
	}

	for _, user := range userList {
		fmt.Fprintf(w, " %v :: %v\n", user.UserID, user.Username)
	}
}

func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func initRouter(db *database.Database) *http.ServeMux {
	err := db.Ping()
	if err != nil {
		log.Fatalf("initRouter: %v\n", err)
	}
	router := http.NewServeMux()

	// -- REGISTER ALL ROUTES HERE --
	router.HandleFunc("/", wrapWithDB(db, handleHome))
	router.HandleFunc("/users", wrapWithDB(db, handleUsers))

	return router
}

package main

import (
	"budget-helper/database"
	"budget-helper/models"
	"budget-helper/user"
	"context"
	"fmt"
	"log"
	"net/http"
)

func home(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Panic(err)
	}
}

func users(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Fetching users...\n")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	repo := user.NewUserRepo(db)
	repo.GetAllUsers()
}

func wrapWithDB(db *database.Database, f func(*database.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func initRouter() *http.ServeMux {
	// -- INITIALIZE DATABASE AND ROUTES --
	db := database.NewDatabase()
	err := db.Ping()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// TODO: This works, but when I try to do this in the UserRepo, I get
	// an error saying the DB is closed.
	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	for idx, val := range users {
		fmt.Printf("%v :: %v\n", idx, val)
	}

	router := http.NewServeMux()
	// ------------------------------------

	// REGISTER ALL ROUTES HERE
	router.HandleFunc("/", wrapWithDB(db, home))
	//router.HandleFunc("/users", wrapWithDB(db, users))

	defer db.Close()
	return router
}

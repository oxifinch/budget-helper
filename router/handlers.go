package router

import (
	"budget-helper/database"
	"budget-helper/users"
	"fmt"
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
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
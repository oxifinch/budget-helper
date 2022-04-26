package router

import (
	"fmt"
	"log"
	"net/http"
)

func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}
}

func (rt *Router) handleUsers(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Fetching users...\n")
	if err != nil {
		log.Fatalf("handleUsers: %v\n", err)
	}

	_, err = rt.UserRepo.GetAllUsers()
	if err != nil {
		log.Fatalf("GetAllUsers: %v\n", err)
	}

	//for _, user := range userList {
	//	fmt.Fprintf(w, " %v :: %v\n", user.UserID, user.Username)
	//}
}

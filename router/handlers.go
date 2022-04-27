package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tmplHome  = template.Must(template.ParseFiles("./templates/pages/home.html"))
	tmplLogin = template.Must(template.ParseFiles("./templates/pages/login.html"))
)

func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	err := tmplHome.Execute(w, nil)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}
}

func (rt *Router) handleLogin(w http.ResponseWriter, r *http.Request) {
	err := tmplLogin.Execute(w, nil)
	if err != nil {
		log.Fatalf("handleLogin: %v\n", err)
	}
}

func (rt *Router) handleUsers(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Fetching users...\n")
	if err != nil {
		log.Fatalf("handleUsers: %v\n", err)
	}

	userList, err := rt.UserRepo.GetAll()
	if err != nil {
		log.Fatalf("handleUsers: %v\n", err)
	}

	for _, user := range userList {
		fmt.Fprintf(w, " %v :: %v\n", user.UserID, user.Username)
	}
}

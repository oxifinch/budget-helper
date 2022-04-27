package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageSettings struct {
	AppTitle  string
	PageTitle string
	filename  string
}

func (p *PageSettings) GetPagePath() string {
	return fmt.Sprintf("./templates/pages/%v", p.filename)
}

func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	data := PageSettings{
		AppTitle:  "Budget Helper",
		PageTitle: "Home",
		filename:  "home.html",
	}

	path := data.GetPagePath()
	t, err := template.New(data.filename).ParseFiles(path)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
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

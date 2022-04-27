package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageSettings struct {
	PageTitle string
	filename  string
	filepath  string
}

func NewPageSettings(pageTitle string, filename string) *PageSettings {
	filepath := fmt.Sprintf("./templates/pages/%v", filename)

	return &PageSettings{
		PageTitle: pageTitle,
		filename:  filename,
		filepath:  filepath,
	}
}

func CreateTemplate(p *PageSettings) (*template.Template, error) {
	return template.New(p.filename).ParseFiles(p.filepath)
}

func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {

	p := NewPageSettings("Home", "home.html")

	t, err := CreateTemplate(p)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}

	err = t.Execute(w, p)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}
}

func (rt *Router) handleLogin(w http.ResponseWriter, r *http.Request) {
	p := NewPageSettings("Sign in", "login.html")

	t, err := CreateTemplate(p)
	if err != nil {
		log.Fatalf("handleLogin: %v\n", err)
	}

	err = t.Execute(w, p)
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

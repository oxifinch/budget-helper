package router

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmplHome  = template.Must(template.ParseFiles("./templates/pages/home.html", "./templates/base.html"))
	tmplLogin = template.Must(template.ParseFiles("./templates/pages/login.html", "./templates/base.html"))
)

func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  "Budget Helper",
		PageTitle: "Home",
	}

	err := tmplHome.ExecuteTemplate(w, "base", data)
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

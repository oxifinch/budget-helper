package router

import (
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

package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	AppTitle = "Budget Helper"
)

var (
	tmplHome      = template.Must(template.ParseFiles("./templates/pages/home.html", "./templates/base.html"))
	tmplError     = template.Must(template.ParseFiles("./templates/pages/error.html", "./templates/base.html"))
	tmplLogin     = template.Must(template.ParseFiles("./templates/pages/login.html", "./templates/base.html"))
	tmplRegister  = template.Must(template.ParseFiles("./templates/pages/register.html", "./templates/base.html"))
	tmplDashboard = template.Must(template.ParseFiles("./templates/pages/dashboard.html", "./templates/base.html"))
)

func displayErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, statusMessage string, detailedMessage string) {
	data := struct {
		AppTitle        string
		PageTitle       string
		StatusCode      int
		StatusMessage   string
		DetailedMessage string
	}{
		AppTitle:        AppTitle,
		PageTitle:       fmt.Sprint(statusCode),
		StatusCode:      statusCode,
		StatusMessage:   statusMessage,
		DetailedMessage: detailedMessage,
	}

	w.WriteHeader(statusCode)
	err := tmplError.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("displayErrorPage: %v\n", err)
	}
}

func methodValid(r *http.Request, methodName string) bool {
	return r.Method == strings.ToUpper(methodName)
}

func nameAndPassValid(username string, password string) bool {
	return strings.TrimSpace(username) != "" && strings.TrimSpace(password) != ""
}

// -- HOME, ABOUT & MISC PAGES --
func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Home",
	}

	err := tmplHome.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleHome: %v\n", err)
	}
}

// -- USERS & AUTHENTICATION --
func (rt *Router) handleLogin(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Login",
	}

	err := tmplLogin.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleLogin: %v\n", err)
	}
}

func (rt *Router) handleLoginSave(w http.ResponseWriter, r *http.Request) {
	if !methodValid(r, "POST") {
		displayErrorPage(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "The resource you requested does not support the method used.")
	}

	username := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request", "One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.GetByCredentials(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound, "Not Found", "We found no user with the provided credentials in the database. Please check your username and password, and try again.")
	}

	// TODO: Set session before redirecting
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (rt *Router) handleRegister(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Register",
	}

	err := tmplRegister.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleRegister: %v\n", err)
	}
}

func (rt *Router) handleRegisterSave(w http.ResponseWriter, r *http.Request) {
	if !methodValid(r, "POST") {
		displayErrorPage(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "The resource you requested does not support the method used.")
	}

	username := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request", "One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.Create(username, password)
	if err != nil {
		log.Fatalf("handleRegisterSave: %v\n", err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Dashboard",
	}

	err := tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
}

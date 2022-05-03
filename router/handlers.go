package router

import (
	"budget-helper/database"
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
	tmplHome      = addTemplate("pages/home.html")
	tmplError     = addTemplate("pages/error.html")
	tmplLogin     = addTemplate("pages/login.html")
	tmplRegister  = addTemplate("pages/register.html")
	tmplDashboard = addTemplate("pages/dashboard.html")
)

const (
	GET  = "GET"
	POST = "POST"
)

func addTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path), "./templates/base.html"))
}

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
	if r.Method != POST {
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
	if r.Method != POST {
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
		Budget    *database.Budget
	}{
		AppTitle:  AppTitle,
		PageTitle: "Dashboard",
	}

	// TODO: Check for authentication and give user the right dashboard.
	// Use user 1 and dashboard 1 for now.
	b, err := rt.BudgetRepo.Get(1)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
	data.Budget = b

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
}

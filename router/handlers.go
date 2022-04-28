package router

import (
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

func validateMethod(w http.ResponseWriter, r *http.Request, methodName string) {
	if r.Method == strings.ToUpper(methodName) {
		return
	}

	data := struct {
		AppTitle        string
		PageTitle       string
		StatusCode      int
		StatusMessage   string
		DetailedMessage string
	}{
		AppTitle:        AppTitle,
		PageTitle:       "Error: 405",
		StatusCode:      http.StatusMethodNotAllowed,
		StatusMessage:   "Method Not Allowed",
		DetailedMessage: "The resource you requested does not support the method used. Please try again by submitting a POST request.",
	}

	w.WriteHeader(data.StatusCode)
	err := tmplError.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("validateMethod: %v\n", err)
	}
}

func validateUsernameAndPassword(w http.ResponseWriter, username string, password string) {
	if strings.TrimSpace(username) != "" && strings.TrimSpace(password) != "" {
		return
	}

	data := struct {
		AppTitle        string
		PageTitle       string
		StatusCode      int
		StatusMessage   string
		DetailedMessage string
	}{
		AppTitle:        AppTitle,
		PageTitle:       "Error: 400",
		StatusCode:      http.StatusBadRequest,
		StatusMessage:   "Bad Request",
		DetailedMessage: "One or more fields was not submitted. Please try again.",
	}

	w.WriteHeader(data.StatusCode)
	err := tmplError.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("validateUsernameAndPassword: %v\n", err)
	}
}

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
	validateMethod(w, r, "POST")

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
	validateMethod(w, r, "POST")

	username := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	validateUsernameAndPassword(w, username, password)

	_, err := rt.UserRepo.Create(username, password)
	if err != nil {
		log.Fatalf("handleRegisterSave: %v\n", err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

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

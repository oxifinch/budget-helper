package router

import (
	"log"
	"net/http"
	"strings"
)

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
		displayErrorPage(w, r, http.StatusMethodNotAllowed, "Method Not Allowed",
			"The resource you requested does not support the method used.")
	}

	username := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request",
			"One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.GetByCredentials(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound, "Not Found",
			"We found no user with the provided credentials in the database. Please check your username and password, and try again.")
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
		displayErrorPage(w, r, http.StatusMethodNotAllowed, "Method Not Allowed",
			"The resource you requested does not support the method used.")
	}

	username := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request",
			"One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.Create(username, password)
	if err != nil {
		log.Fatalf("handleRegisterSave: %v\n", err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

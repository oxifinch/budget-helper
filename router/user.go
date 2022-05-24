package router

import (
	"budget-helper/database"
	"net/http"
	"strconv"
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
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) handleLoginSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	username := trimmedFormValue(r, "username")
	password := trimmedFormValue(r, "password")
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.GetByCredentials(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"We found no user with the provided credentials in the database. Please check your username and password, and try again.")
	}

	// TODO: Set session before redirecting and get user's actual active budget
	http.Redirect(w, r, "/dashboard?id=1", http.StatusSeeOther)
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
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) handleRegisterSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	username := trimmedFormValue(r, "username")
	password := trimmedFormValue(r, "password")
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
	}

	_, err := rt.UserRepo.Create(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created.. Please try again later.")
	}

	http.Redirect(w, r, "/newBudget", http.StatusSeeOther)
}

func (rt *Router) handleSettings(w http.ResponseWriter, r *http.Request) {
	// TODO: Check that user is auhenticated. Use UserID 1 for now.
	id := 1

	user, err := rt.UserRepo.Get(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	data := struct {
		AppTitle  string
		PageTitle string
		User      *database.User
	}{
		AppTitle:  AppTitle,
		PageTitle: "Settings",
		User:      user,
	}

	err = tmplSettings.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

}

func (rt *Router) handleSettingsAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	err = tmplPartSettingsAccount.Execute(w, nil)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) handleSettingsIncomeExpenses(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(uint(id))
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you requested could not be found. Check the request and try again.")
	}

	data := struct {
		ID             uint
		IncomeExpenses []database.IncomeExpense
	}{
		ID:             uint(id),
		IncomeExpenses: ies,
	}

	err = tmplPartSettingsIncomeExpenses.Execute(w, data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

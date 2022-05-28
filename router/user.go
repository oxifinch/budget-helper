package router

import (
	"budget-helper/database"
	"net/http"
)

// -- USERS & AUTHENTICATION --
func (rt *Router) handleLogin(w http.ResponseWriter, r *http.Request) {
	if rt.userIsLoggedIn(w, r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

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
	if rt.userIsLoggedIn(w, r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

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

	u, err := rt.UserRepo.GetByCredentials(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"We found no user with the provided credentials in the database. Please check your username and password, and try again.")
	}

	session, err := rt.Store.Get(r, "session")
	// fmt.Print("Session: %v\n", session)
	session.Values["userID"] = u.ID
	session.Save(r, w)

	// TODO: Set session before redirecting and get user's actual active budget
	http.Redirect(w, r, "/dashboard?id=1", http.StatusSeeOther)
}

func (rt *Router) handleRegister(w http.ResponseWriter, r *http.Request) {
	if rt.userIsLoggedIn(w, r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

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
	if rt.userIsLoggedIn(w, r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

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

	uID, err := rt.UserRepo.Create(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created.. Please try again later.")
	}

	session, err := rt.Store.Get(r, "session")
	session.Values["userID"] = uID
	session.Save(r, w)

	http.Redirect(w, r, "/newBudget", http.StatusSeeOther)
}

func (rt *Router) handleSettings(w http.ResponseWriter, r *http.Request) {
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	id, err := rt.getUserIDFromSession(r)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to handle your user session. Please try again later.")
		return
	}

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
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	err := tmplPartSettingsAccount.Execute(w, nil)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) handleSettingsIncomeExpenses(w http.ResponseWriter, r *http.Request) {
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	id, err := rt.getUserIDFromSession(r)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(id)
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

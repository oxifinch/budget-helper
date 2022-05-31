package router

import (
	"budget-helper/auth"
	"budget-helper/database"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// -- USERS & AUTHENTICATION --
func (rt *Router) handleLogin(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if found {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
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
		return
	}
}

func (rt *Router) handleLoginSave(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if found {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	username := trimmedFormValue(r, "username")
	password := trimmedFormValue(r, "password")
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
		return
	}

	u, err := rt.UserRepo.GetByCredentials(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"We found no user with the provided credentials in the database. Please check your username and password, and try again.")
		return
	}

	session, err := rt.Store.Get(r, "session")
	session.Values["userID"] = u.ID
	session.Save(r, w)

	// TODO: Set session before redirecting and get user's actual active budget
	http.Redirect(w, r, "/dashboard?id=1", http.StatusSeeOther)
}

func (rt *Router) handleRegister(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if found {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
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
		return
	}
}

func (rt *Router) handleRegisterSave(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if found {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	username := trimmedFormValue(r, "username")
	password := trimmedFormValue(r, "password")
	if !nameAndPassValid(username, password) {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
		return
	}

	uID, err := rt.UserRepo.Create(username, password)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created.. Please try again later.")
		return
	}

	session, err := rt.Store.Get(r, "session")
	session.Values["userID"] = uID
	session.Save(r, w)

	http.Redirect(w, r, "/newBudget", http.StatusSeeOther)
}

func (rt *Router) handleSettings(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	user, err := rt.UserRepo.Get(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	data := struct {
		AppTitle  string
		PageTitle string
		User      *database.User
		Currency  string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Settings",
		User:      user,
	}

	err = tmplSettings.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

}

func (rt *Router) handleSettingsAccount(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	user, err := rt.UserRepo.Get(uint(id))
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to get your user information. Please try again later.")
		return
	}

	data := struct {
		User          *database.User
		Currency      string
		Budget        *database.Budget
		BudgetExpired bool
	}{
		User: user,
	}

	switch user.Currency {
	case database.USD:
		data.Currency = "USD"
		break
	case database.EUR:
		data.Currency = "EUR"
		break
	case database.SEK:
		data.Currency = "SEK"
		break
	default:
		data.Currency = "Unknown"
		break
	}

	// Getting the budget's end date and comparing it to the current date, to notify
	// the user if their current budget period has expired. If so, they should be given
	// the option to create a new one.
	b, err := rt.BudgetRepo.GetInfo(user.ActiveBudgetID)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to process your request. Please try again later.")
		return
	}
	data.Budget = b

	curDate := time.Now()
	endDate, err := time.Parse("2006-01-02", b.EndDate)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to process your request. Please try again later.")
		return
	}
	data.BudgetExpired = curDate.After(endDate)

	err = tmplPartSettingsAccount.Execute(w, data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

func (rt *Router) handleSettingsIncomeExpenses(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you requested could not be found. Check the request and try again.")
		return
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
		return
	}
}

func (rt *Router) handleSettingsSaveAccount(w http.ResponseWriter, r *http.Request) {
	// TODO: Check for user authentication. Use ID 1 for now.
	id := uint(1)

	// TODO: Is this the correct way of closing requests? Using return here to
	// exit out of the function since there is nothing more to do.
	defer r.Body.Close()
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	// Validate POST values
	err := r.ParseForm()
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to process your request. Please try again later.")
		return
	}

	postCurrency := strings.TrimSpace(r.PostFormValue("currency"))
	if postCurrency == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted in your request. Please check the request and try again.")
		return
	}

	selectedCurrency, err := strconv.Atoi(postCurrency)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to process your request. Please try again later.")
		return
	}

	// TODO: There is probably a better way to do this but I'm having brainfarts
	// at the time of writing. Come back and revise this.
	var currency database.Currency
	switch selectedCurrency {
	case 1:
		currency = database.USD
		break
	case 2:
		currency = database.EUR
		break
	case 3:
		currency = database.SEK
		break
	}

	// TODO: Get user's active budget ID. Use ID 1 for now.
	err = rt.UserRepo.UpdateSettings(id, 1, currency)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Your settings could not be saved at this time. Please try again later.")
		return
	}

	err = tmplPartAccountConfirmed.Execute(w, nil)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

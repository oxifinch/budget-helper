package router

import (
	"budget-helper/auth"
	"budget-helper/database"
	"log"
	"net/http"
	"strconv"
)

func (rt *Router) handleSettingsDataIncomeExpenses(w http.ResponseWriter, r *http.Request) {
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

	err = tmplPartSettingsDataIncomeExpenses.Execute(w, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

}

func (rt *Router) handleIncomeExpensesCreate(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	err := r.ParseForm()
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	// Validate POST values.
	postLabel := trimmedFormValue(r, "label")
	postDay := trimmedFormValue(r, "day")
	postAmount := trimmedFormValue(r, "amount")

	if postLabel == "" || postDay == "" || postAmount == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
		return
	}

	day, err := strconv.Atoi(postDay)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	amount, err := strconv.ParseFloat(postAmount, 64)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	_, err = rt.IncomeExpenseRepo.Create(id, postLabel, uint(day), amount)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created. Please try again later.")
		return
	}

	data := struct {
		ID     uint
		Create bool
		Update bool
		Delete bool
	}{
		ID:     uint(id),
		Create: true,
		Update: false,
		Delete: false,
	}

	err = tmplPartIncomeExpenseConfirmed.Execute(w, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

}

func (rt *Router) handleIncomeExpensesUpdate(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	// Validate POST values.
	postID := trimmedFormValue(r, "id")
	postLabel := trimmedFormValue(r, "label")
	postDay := trimmedFormValue(r, "day")
	postAmount := trimmedFormValue(r, "amount")
	postEnabled := trimmedFormValue(r, "enabled")

	if postID == "" || postLabel == "" || postDay == "" || postAmount == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
		return
	}

	enabled := true
	if postEnabled == "" {
		enabled = false
	}

	// Parse numerical values and create copies with the correct type.
	ieID, err := strconv.Atoi(postID)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
	if ieID < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The resource ID submitted with the request is invalid. Double-check the request and try again.")
		return
	}

	day, err := strconv.Atoi(postDay)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	amount, err := strconv.ParseFloat(postAmount, 64)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	err = rt.IncomeExpenseRepo.Update(uint(ieID), postLabel, uint(day), amount, enabled)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	data := struct {
		ID     uint
		Create bool
		Update bool
		Delete bool
	}{
		ID:     id,
		Create: false,
		Update: true,
		Delete: false,
	}

	err = tmplPartIncomeExpenseConfirmed.Execute(w, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

func (rt *Router) handleIncomeExpensesDelete(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	if r.Method != DELETE {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	queryID := r.URL.Query().Get("id")
	if queryID == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"No resource ID was submitted in the request. Check the URL and try again.")
		return
	}

	ieID, err := strconv.Atoi(queryID)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The resource ID submitted in the reqest is invalid. Check the URL and try again.")
		return
	}

	err = rt.IncomeExpenseRepo.Delete(uint(ieID))
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be deleted. Please try again later.")
		return
	}

	// TODO: Load the actual UserID here instead of the default 1.
	data := struct {
		ID     uint
		Create bool
		Update bool
		Delete bool
	}{
		ID:     1,
		Create: false,
		Update: false,
		Delete: true,
	}

	err = tmplPartIncomeExpenseConfirmed.Execute(w, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

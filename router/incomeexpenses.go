package router

import (
	"budget-helper/database"
	"log"
	"net/http"
	"strconv"
)

func (rt *Router) handleSettingsDataIncomeExpenses(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	if queryID == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request did not include a resource ID. Check the URL and try again.")
	}

	id, err := strconv.Atoi(queryID)
	if err != nil {
		log.Fatalf("handleSettingsDataIncomeExpenses: %v\n", err)
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(uint(id))
	if err != nil {
		log.Fatalf("handleSettingsDataIncomeExpenses: %v\n", err)
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
		log.Fatalf("handleSettingsDataIncomeExpenses: %v\n", err)
	}

}

func (rt *Router) handleIncomeExpensesCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
	}

	// Validate POST values.
	postID := trimmedFormValue(r, "id")
	postLabel := trimmedFormValue(r, "label")
	postDay := trimmedFormValue(r, "day")
	postAmount := trimmedFormValue(r, "amount")

	if postID == "" || postLabel == "" || postDay == "" || postAmount == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"One or more fields was not submitted. Please try again.")
	}

	// Parse numerical values and create copies with the correct type.
	id, err := strconv.Atoi(postID)
	if err != nil {
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
	}
	day, err := strconv.Atoi(postDay)
	if err != nil {
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
	}

	amount, err := strconv.ParseFloat(postAmount, 64)
	if err != nil {
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
	}

	_, err = rt.IncomeExpenseRepo.Create(uint(id), postLabel, uint(day), amount)
	if err != nil {
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
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
		log.Fatalf("handleIncomeExpensesCreate: %v\n", err)
	}

}

func (rt *Router) handleIncomeExpensesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	// TODO: Make sure the user is logged in and get their ID.

	err := r.ParseForm()
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
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
	}

	// If "enabled" was not checked(not included), default to false value.
	enabled := true
	if postEnabled == "" {
		enabled = false
	}

	// Parse numerical values and create copies with the correct type.
	id, err := strconv.Atoi(postID)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}
	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The resource ID submitted with the request is invalid. Double-check the request and try again.")
	}

	day, err := strconv.Atoi(postDay)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}

	amount, err := strconv.ParseFloat(postAmount, 64)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}

	err = rt.IncomeExpenseRepo.Update(uint(id), postLabel, uint(day), amount, enabled)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
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
		Update: true,
		Delete: false,
	}

	err = tmplPartIncomeExpenseConfirmed.Execute(w, data)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}
}

func (rt *Router) handleIncomeExpensesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != DELETE {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	// TODO: Make sure the user is logged in and get their ID.

	queryID := r.URL.Query().Get("id")
	if queryID == "" {
		displayErrorPage(w, r, http.StatusBadRequest,
			"No resource ID was submitted in the request. Check the URL and try again.")
	}

	id, err := strconv.Atoi(queryID)
	if err != nil {
		log.Fatalf("handleIncomeExpensesDelete: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The resource ID submitted in the reqest is invalid. Check the URL and try again.")
	}

	err = rt.IncomeExpenseRepo.Delete(uint(id))
	if err != nil {
		log.Fatalf("handleIncomeExpensesDelete: %v\n", err)
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
		log.Fatalf("handleIncomeExpensesDelete: %v\n", err)
	}
}

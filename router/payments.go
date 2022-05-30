package router

import (
	"budget-helper/auth"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ----------------------------------------------------------------------------------
// Payments can be retrieved from several sources. All payments from a specific User,
// a single Budget, a single BudgetCategory, or all payments across a global Category
// on a User account.
// ----------------------------------------------------------------------------------
// Get all payments from a specific Budget.
func (rt *Router) handlePaymentsBudget(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
	}

	bID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	if bID < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	ps, err := rt.PaymentRepo.GetAllByBudgetID(uint(bID))
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you requested could not be found. Check the request and try again.")
	}

	err = tmplPartPayment.Execute(w, ps)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

}

// Get all payments from a BudgetCategory.
func (rt *Router) handlePaymentsBudgetCategory(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
	}

	bcID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	if bcID < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	ps, err := rt.PaymentRepo.GetAllByBudgetCategoryID(uint(bcID))
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you request could not be found. Check the request and try again.")
	}

	err = tmplPartPayment.Execute(w, ps)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

}

// Get all payments from a User's category across all their Budgets.
func (rt *Router) handlePaymentsCategory(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
	}

	cID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	if cID < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	ps, err := rt.PaymentRepo.GetAllByCategoryID(uint(cID))
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you request could not be found. Check the request and try again.")
	}

	err = tmplPartPayment.Execute(w, ps)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

// Get all payments from the User account.
func (rt *Router) handlePaymentsUser(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
	}

	ps, err := rt.PaymentRepo.GetAllByUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusNotFound,
			"The resource you requested could not be found. Check the request and try again.")
	}

	err = tmplPartPayment.Execute(w, ps)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

}

func (rt *Router) handlePaymentSave(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
	}

	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	// Validate POST contents
	bcID := trimmedFormValue(r, "budget_category_id")
	date := trimmedFormValue(r, "date")
	amount := trimmedFormValue(r, "amount")
	note := trimmedFormValue(r, "note")

	// Convert ID and amount
	bcIDInt, err := strconv.Atoi(bcID)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	data := struct {
		Success      bool
		ErrorMessage string
	}{
		Success: true,
	}

	_, err = rt.PaymentRepo.Create(date, amountFloat, note, uint(bcIDInt))
	if err != nil {
		data.Success = false
		data.ErrorMessage = fmt.Sprintf("%v", err)
	}

	err = tmplPartPaymentConfirmed.Execute(w, data)
	if err != nil {
		log.Printf("error: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

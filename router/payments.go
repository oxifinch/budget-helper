package router

import (
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	Payments, err := rt.PaymentRepo.GetAllByBudgetID(uint(id))
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

	err = tmplPartPayment.Execute(w, Payments)
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

}

// Get all payments from a BudgetCategory.
func (rt *Router) handlePaymentsBudgetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalf("handlePaymentsBudgetCategory: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	Payments, err := rt.PaymentRepo.GetAllByBudgetCategoryID(uint(id))
	if err != nil {
		log.Fatalf("handlePaymentsBudgetCategory: %v\n", err)
	}

	err = tmplPartPayment.Execute(w, Payments)
	if err != nil {
		log.Fatalf("handlePaymentsBudgetCategory: %v\n", err)
	}

}

// Get all payments from a User's category across all their Budgets.
func (rt *Router) handlePaymentsCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalf("handlePaymentsCategory: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	Payments, err := rt.PaymentRepo.GetAllByCategoryID(uint(id))
	// if err != nil {
	// 	log.Fatalf("handlePaymentsCategory: %v\n", err)
	// }

	err = tmplPartPayment.Execute(w, Payments)
	if err != nil {
		log.Fatalf("handlePaymentsCategory: %v\n", err)
	}

}

// Get all payments from the User account.
func (rt *Router) handlePaymentsUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalf("handlePaymentsUser: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	Payments, err := rt.PaymentRepo.GetAllByUserID(uint(id))
	if err != nil {
		log.Fatalf("handlePaymentsUser: %v\n", err)
	}

	err = tmplPartPayment.Execute(w, Payments)
	if err != nil {
		log.Fatalf("handlePaymentsUser: %v\n", err)
	}

}

func (rt *Router) handlePaymentSave(w http.ResponseWriter, r *http.Request) {
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
		log.Fatalf("handlePaymentSave: %v\n", err)
	}
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatalf("handlePaymentSave: %v\n", err)
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
		log.Fatalf("handlePaymentSave: %v\n", err)
	}
}

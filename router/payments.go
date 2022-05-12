package router

import (
	"budget-helper/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (rt *Router) handlePaymentsBudget(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The request included an invalid resource ID. Check the URL and try again.")
	}

	ps, err := rt.PaymentRepo.GetAllByBudgetID(uint(id))
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

	data := struct {
		Payments []database.Payment
	}{
		Payments: ps,
	}

	err = tmplPartPayment.Execute(w, data)
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
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

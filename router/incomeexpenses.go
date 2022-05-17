package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (rt *Router) handleIncomeExpensesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate", err)
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

	fmt.Printf("UPDATING IncomeExpense with ID: %v\n", id)
	err = rt.IncomeExpenseRepo.Update(uint(id), postLabel, uint(day), amount)
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}

	// Send just a checkbox or something similar, no need for a full template
	_, err = fmt.Fprintf(w, "SAVED")
	if err != nil {
		log.Fatalf("handleIncomeExpensesUpdate: %v\n", err)
	}
}

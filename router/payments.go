package router

import (
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
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request", "The request included an invalid resource ID. Check the URL and try again.")
	}

	ps, err := rt.PaymentRepo.GetAllByBudgetID(uint(id))
	if err != nil {
		log.Fatalf("main: %v\n", err)
	}

	for _, p := range *ps {
		fmt.Fprintf(w, "%v :: Date: %v | Amount: %v\n", p.ID, p.Date, p.Amount)
		fmt.Fprintf(w, "\t - Belongs to BudgetCategory: %v\n", p.BudgetCategory.Category.Name)
	}
}

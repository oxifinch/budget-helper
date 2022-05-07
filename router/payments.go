package router

import (
	"budget-helper/database"
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

	data := struct {
		Payments []database.Payment
	}{
		Payments: *ps,
	}

	err = tmplPartPayment.Execute(w, data)
	if err != nil {
		log.Fatalf("handlePaymentsBudget: %v\n", err)
	}

}

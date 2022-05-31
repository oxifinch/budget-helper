package router

import (
	"budget-helper/auth"
	"budget-helper/database"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	data := struct {
		AppTitle         string
		PageTitle        string
		Budget           *database.Budget
		Categories       []database.Category
		BalanceRemaining string
		BufferRemaining  string
		PercentageSpent  int
	}{
		AppTitle:  AppTitle,
		PageTitle: "Dashboard",
	}

	b, err := rt.BudgetRepo.GetByUserID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			displayErrorPage(w, r, http.StatusNotFound,
				"The resource you requested could not be found. Check the request and try again.")
			return
		} else {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"Something went wrong. Please try again later.")
			return
		}
	}
	data.Budget = b

	bcAllocated := budgetCategoriesAllocated(b)
	bcSpent := budgetCategoriesSpent(b)

	bufAllocated := budgetBufferAllocated(b)
	bufSpent := budgetBufferSpent(b)

	data.BalanceRemaining = fmt.Sprintf("%.2f", bcAllocated-bcSpent)
	data.BufferRemaining = fmt.Sprintf("%.2f", bufAllocated-bufSpent)

	if bcSpent > 0 {
		data.PercentageSpent = budgetPercentageSpent(b)
	}

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

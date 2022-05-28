package router

import (
	"budget-helper/database"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	session, err := rt.Store.Get(r, "session")
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to handle your user session. Please try again later.")
		return
	}

	id, isset := session.Values["userID"]
	if !isset {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to handle your user session. Please try again later.")
		return
	}

	userID, ok := id.(uint)
	if !ok {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to handle your user session. Please try again later.")
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

	b, err := rt.BudgetRepo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			displayErrorPage(w, r, http.StatusNotFound,
				"The resource you requested could not be found. Check the request and try again.")
		} else {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"Something went wrong. Please try again later.")
		}
	}
	data.Budget = b

	var totalSpent float64
	var totalAllocated float64
	for _, bc := range b.BudgetCategories {
		totalAllocated += bc.Allocated

		for _, p := range bc.Payments {
			totalSpent += p.Amount
		}
	}
	if totalSpent > 0 {
		data.PercentageSpent = int((totalSpent / totalAllocated) * 100)
	}
	data.BalanceRemaining = fmt.Sprintf("%.2f", (totalAllocated - totalSpent))
	// TODO: BufferRemaining should be calculated here.
	// BufferRemaining = (Budget.Allocated - BudgetCategories.Allocated) - (Uncategorized payments + BudgetCategory deficits)

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

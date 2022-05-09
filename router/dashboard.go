package router

import (
	"budget-helper/database"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle        string
		PageTitle       string
		Budget          *database.Budget
		Categories      []database.Category
		BudgetRemaining float32
		BufferRemaining float32
		PercentageSpent int
	}{
		AppTitle:  AppTitle,
		PageTitle: "Dashboard",
	}

	getBudgetID := strings.TrimSpace(r.URL.Query().Get("id"))
	id, err := strconv.Atoi(getBudgetID)
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request",
			"The ID of the resource you are trying to access was not included in the request. Check the URL and try again.")
	}
	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest, "Bad Request",
			"The ID submitted in the request is invalid. Check the URL and try again.")
	}

	// TODO: Check for authentication and give user the right dashboard.
	log.Printf("Looking for Budget with ID: %v...\n", id)
	b, err := rt.BudgetRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			displayErrorPage(w, r, http.StatusNotFound, "Not Found",
				"The resource you requested could not be found in our database. Check the request and try again.")
		} else {
			log.Fatalf("handleDashboard: %v\n", err)
		}
	}
	data.Budget = b

	var totalSpent float32
	var totalAllocated float32
	for _, bc := range b.BudgetCategories {
		totalAllocated += bc.Allocated

		for _, p := range bc.Payments {
			totalSpent += p.Amount
		}
	}
	data.PercentageSpent = int((totalSpent / totalAllocated) * 100)
	data.BudgetRemaining = totalAllocated - totalSpent
	// TODO: BufferRemaining should be calculated here.
	// BufferRemaining = (Budget.Allocated - BudgetCategories.Allocated) - (Uncategorized payments + BudgetCategory deficits)

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
}

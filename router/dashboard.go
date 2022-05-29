package router

import (
	"budget-helper/database"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
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

	getBudgetID := strings.TrimSpace(r.URL.Query().Get("id"))
	id, err := strconv.Atoi(getBudgetID)
	if err != nil {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The ID of the resource you are trying to access was not included in the request. Check the URL and try again.")
	}
	if id < 1 {
		displayErrorPage(w, r, http.StatusBadRequest,
			"The ID submitted in the request is invalid. Check the URL and try again.")
	}

	// TODO: Check for authentication and give user the right dashboard.
	log.Printf("Looking for Budget with ID: %v...\n", id)
	b, err := rt.BudgetRepo.Get(id)
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
	var bcAllocated float64
	var bufferSpent float64
	for _, bc := range b.BudgetCategories {
		bcAllocated += bc.Allocated

		var spentInBC float64
		for _, p := range bc.Payments {
			spentInBC += p.Amount
		}
		totalSpent += spentInBC

		if spentInBC > bc.Allocated {
			bufferSpent += (bc.Allocated - spentInBC)
		}
	}

	if totalSpent > 0 {
		data.PercentageSpent = int((totalSpent / bcAllocated) * 100)
	}
	data.BalanceRemaining = fmt.Sprintf("%.2f", (bcAllocated - totalSpent))

	data.BufferRemaining = fmt.Sprintf("%.2f", (b.Allocated-bcAllocated)+bufferSpent)

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

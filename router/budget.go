package router

import (
	"budget-helper/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (rt *Router) handleNewBudget(w http.ResponseWriter, r *http.Request) {
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	id, err := rt.getUserIDFromSession(r)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	// TODO: Get user ID above
	categories, err := rt.CategoryRepo.GetAllWithUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	data := struct {
		AppTitle         string
		PageTitle        string
		Categories       []database.Category
		DefaultAllocated float64
	}{
		AppTitle:   AppTitle,
		PageTitle:  "New Budget",
		Categories: categories,
	}

	for _, ie := range ies {
		if !ie.Enabled {
			continue
		}
		data.DefaultAllocated += ie.Amount
	}

	err = tmplNewBudget.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) handleNewBudgetSave(w http.ResponseWriter, r *http.Request) {
	if !rt.userIsLoggedIn(w, r) {
		displayLoginRequired(w, r)
		return
	}

	// Validate POST
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
	}

	postAllocated := trimmedFormValue(r, "allocated")
	postStartDate := trimmedFormValue(r, "start_date")
	postEndDate := trimmedFormValue(r, "end_date")

	// Check and convert POST values
	allocated, err := strconv.ParseFloat(postAllocated, 32)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	_, err = time.Parse("2006-01-02", postStartDate)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	_, err = time.Parse("2006-01-02", postEndDate)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")

	}

	// TODO: Should a HTTP status 201 be sent here?
	budgetID, err := rt.BudgetRepo.Create(postStartDate, postEndDate, allocated)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created.. Please try again later.")
	}

	// Loop through categories and create new BudgetCategories
	for key := range r.PostForm {
		if !strings.Contains(key, "bc_allocated_") {
			continue
		}

		var ctID uint
		_, err := fmt.Sscanf(key, "bc_allocated_%d", &ctID)
		if err != nil {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"Something went wrong. Please try again later.")
		}

		allocated, err = strconv.ParseFloat(r.PostFormValue(key), 32)
		if err != nil {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"Something went wrong. Please try again later.")

		}

		_, err = rt.BudgetRepo.CreateBudgetCategory(budgetID, ctID, allocated)
		if err != nil {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"The resource could not be created.. Please try again later.")

		}
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard?id=%v", budgetID), http.StatusSeeOther)
}

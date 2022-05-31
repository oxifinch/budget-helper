package router

import (
	"budget-helper/auth"
	"budget-helper/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (rt *Router) handleNewBudget(w http.ResponseWriter, r *http.Request) {
	id, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	// TODO: Get user ID above
	categories, err := rt.CategoryRepo.GetAllWithUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(id)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	data := struct {
		AppTitle         string
		PageTitle        string
		Categories       []database.Category
		IncomeExpenses   []database.IncomeExpense
		DefaultAllocated float64
	}{
		AppTitle:       AppTitle,
		PageTitle:      "New Budget",
		Categories:     categories,
		IncomeExpenses: ies,
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
		return
	}
}

func (rt *Router) handleNewBudgetSave(w http.ResponseWriter, r *http.Request) {
	_, found := auth.LoggedInUser(rt.Store, r)
	if !found {
		displayLoginRequired(w, r)
		return
	}

	// Validate POST
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed,
			"The resource you requested does not support the method used.")
		return
	}

	postAllocated := trimmedFormValue(r, "allocated")
	postStartDate := trimmedFormValue(r, "start_date")
	postEndDate := trimmedFormValue(r, "end_date")

	// Check and convert POST values
	allocated, err := strconv.ParseFloat(postAllocated, 32)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	_, err = time.Parse("2006-01-02", postStartDate)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}

	_, err = time.Parse("2006-01-02", postEndDate)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return

	}

	// TODO: Should a HTTP status 201 be sent here?
	budgetID, err := rt.BudgetRepo.Create(postStartDate, postEndDate, allocated)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The resource could not be created.. Please try again later.")
		return
	}

	// TODO: Use actual User ID here. Use 1 for now.
	user, err := rt.UserRepo.Get(1)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"The server was unable to get your user information. Please try again later.")
		return
	}

	err = rt.UserRepo.UpdateSettings(user.ID, budgetID, user.Currency)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Your user settings could not be saved at this time. Please try again later.")
		return
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
			return
		}

		allocated, err = strconv.ParseFloat(r.PostFormValue(key), 32)
		if err != nil {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"Something went wrong. Please try again later.")
			return

		}

		_, err = rt.BudgetRepo.CreateBudgetCategory(budgetID, ctID, allocated)
		if err != nil {
			displayErrorPage(w, r, http.StatusInternalServerError,
				"The resource could not be created.. Please try again later.")
			return

		}
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard?id=%v", budgetID), http.StatusSeeOther)
}

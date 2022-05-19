package router

import (
	"budget-helper/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (rt *Router) handleNewBudget(w http.ResponseWriter, r *http.Request) {
	// TODO: User should be verified before anything else happens.
	// displayLoginRequired(w, r)

	// TODO: Get user ID above
	categories, err := rt.CategoryRepo.GetAllWithUserID(1)
	if err != nil {
		log.Fatalf("handleNewBudget: %v\n", err)
	}

	ies, err := rt.IncomeExpenseRepo.GetAllWithUserID(1)
	if err != nil {
		log.Fatalf("handleNewBudget: %v\n", err)
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
		if ie.Enabled == true {
			data.DefaultAllocated += ie.Amount
		}
	}

	err = tmplNewBudget.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleNewBudget: %v\n", err)
	}
}

func (rt *Router) handleNewBudgetSave(w http.ResponseWriter, r *http.Request) {
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
		log.Fatalf("handleNewBudgetSave: %v\n", err)
	}

	_, err = time.Parse("2006-01-02", postStartDate)
	if err != nil {
		log.Fatalf("handleNewBudgetSave: %v\n", err)
	}

	_, err = time.Parse("2006-01-02", postEndDate)
	if err != nil {
		log.Fatalf("handleNewBudgetSave: %v\n", err)
	}

	// TODO: Should a HTTP status 201 be sent here?
	budgetID, err := rt.BudgetRepo.Create(postStartDate, postEndDate, allocated)
	if err != nil {
		log.Fatalf("handleNewBudgetSave: %v\n", err)
	}

	// Loop through categories and create new BudgetCategories
	for key := range r.PostForm {
		if !strings.Contains(key, "bc_allocated_") {
			continue
		}

		var ctID uint
		_, err := fmt.Sscanf(key, "bc_allocated_%d", &ctID)
		if err != nil {
			log.Fatalf("handleNewBudgetSave: %v\n", err)
		}

		allocated, err = strconv.ParseFloat(r.PostFormValue(key), 32)
		if err != nil {
			log.Fatalf("handleNewBudgetSave: %v\n", err)
		}

		_, err = rt.BudgetRepo.CreateBudgetCategory(budgetID, ctID, allocated)
		if err != nil {
			log.Fatalf("handleNewBudgetSave: %v\n", err)
		}

	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard?id=%v", budgetID), http.StatusSeeOther)
}

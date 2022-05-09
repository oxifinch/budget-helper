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

	data := struct {
		AppTitle   string
		PageTitle  string
		Categories []database.Category
	}{
		AppTitle:   AppTitle,
		PageTitle:  "New Budget",
		Categories: categories,
	}

	err = tmplNewBudget.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleNewBudget: %v\n", err)
	}
}

func (rt *Router) handleNewBudgetSave(w http.ResponseWriter, r *http.Request) {
	// Validate POST
	if r.Method != POST {
		displayErrorPage(w, r, http.StatusMethodNotAllowed, "Method Not Allowed",
			"The resource you requested does not support the method used.")
	}

	postAllocated := strings.TrimSpace(r.PostFormValue("allocated"))
	postStartDate := strings.TrimSpace(r.PostFormValue("start_date"))
	postEndDate := strings.TrimSpace(r.PostFormValue("end_date"))

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
	id, err := rt.BudgetRepo.Create(postStartDate, postEndDate, float32(allocated))
	if err != nil {
		log.Fatalf("handleNewBudgetSave: %v\n", err)
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard?id=%v", id), http.StatusSeeOther)
}

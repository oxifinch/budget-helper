package router

import (
	"budget-helper/database"
	"log"
	"net/http"
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

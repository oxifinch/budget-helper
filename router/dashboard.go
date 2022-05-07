package router

import (
	"budget-helper/database"
	"log"
	"net/http"
)

// -- DASHBOARD & MAIN APP ROUTES --
func (rt *Router) handleDashboard(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle        string
		PageTitle       string
		Budget          *database.Budget
		Categories      []*database.Category
		BudgetRemaining float32
	}{
		AppTitle:  AppTitle,
		PageTitle: "Dashboard",
	}

	// TODO: Check for authentication and give user the right dashboard.
	// Use user 1 and dashboard 1 for now.
	b, err := rt.BudgetRepo.Get(1)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
	data.Budget = b

	// TODO: Display result with only two digits, this is a bit overkill
	// for _, val := range data.Budget.BudgetCategories {
	// 	data.BudgetRemaining = data.BudgetRemaining + val.Remaining
	// }

	err = tmplDashboard.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleDashboard: %v\n", err)
	}
}

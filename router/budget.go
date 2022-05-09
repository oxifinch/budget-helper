package router

import (
	"log"
	"net/http"
)

func (rt *Router) handleNewBudget(w http.ResponseWriter, r *http.Request) {
	// TODO: User should be verified before anything else happens.
	// displayLoginRequired(w, r)

	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "New Budget",
	}

	err := tmplNewBudget.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("handleNewBudget: %v\n", err)
	}
}

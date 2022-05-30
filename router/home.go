package router

import (
	"net/http"
)

// -- HOME, ABOUT & MISC PAGES --
func (rt *Router) handleHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Home",
	}

	err := tmplHome.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
		return
	}
}

package router

import (
	"log"
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
		log.Fatalf("handleHome: %v\n", err)
	}
}

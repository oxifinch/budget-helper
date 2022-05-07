package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	tmplHome      = addTemplate("pages/home.html")
	tmplError     = addTemplate("pages/error.html")
	tmplLogin     = addTemplate("pages/login.html")
	tmplRegister  = addTemplate("pages/register.html")
	tmplDashboard = addTemplate("pages/dashboard.html")
)

const (
	AppTitle = "Budget Helper"
	GET      = "GET"
	POST     = "POST"
)

func addTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path), "./templates/base.html"))
}

func displayErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, statusMessage string, detailedMessage string) {
	data := struct {
		AppTitle        string
		PageTitle       string
		StatusCode      int
		StatusMessage   string
		DetailedMessage string
	}{
		AppTitle:        AppTitle,
		PageTitle:       fmt.Sprint(statusCode),
		StatusCode:      statusCode,
		StatusMessage:   statusMessage,
		DetailedMessage: detailedMessage,
	}

	w.WriteHeader(statusCode)
	err := tmplError.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatalf("displayErrorPage: %v\n", err)
	}
}

func nameAndPassValid(username string, password string) bool {
	return strings.TrimSpace(username) != "" && strings.TrimSpace(password) != ""
}

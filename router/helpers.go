package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// -- TEMPLATES: FULL PAGES --
var (
	tmplHome      = addTemplate("pages/home.html")
	tmplError     = addTemplate("pages/error.html")
	tmplLogin     = addTemplate("pages/login.html")
	tmplRegister  = addTemplate("pages/register.html")
	tmplDashboard = addTemplate("pages/dashboard.html")
)

// -- TEMPLATES: PARTIALS --
var (
	tmplPartPayment          = addPartial("partials/payment.html")
	tmplPartPaymentConfirmed = addPartial("partials/payment-confirmed.html")
)

const (
	AppTitle = "Budget Helper"
	GET      = "GET"
	POST     = "POST"
)

func addTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path), "./templates/base.html"))
}

func addPartial(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path)))
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

func saveJson(data struct{}) error {
	json, err := json.MarshalIndent(data, "", "\t")
	err = os.WriteFile("./data-budget.json", json, 0666)

	return err
}

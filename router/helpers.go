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
	tmplHome          = addTemplate("pages/home.html")
	tmplError         = addTemplate("pages/error.html")
	tmplLoginRequired = addTemplate("pages/loginrequired.html")
	tmplLogin         = addTemplate("pages/login.html")
	tmplRegister      = addTemplate("pages/register.html")
	tmplSettings      = addTemplate("pages/settings.html")
	tmplDashboard     = addTemplate("pages/dashboard.html")
	tmplNewBudget     = addTemplate("pages/newbudget.html")
)

// -- TEMPLATES: PARTIALS --
var (
	tmplPartPayment                    = addPartial("partials/payment.html")
	tmplPartPaymentConfirmed           = addPartial("partials/payment-confirmed.html")
	tmplPartSettingsAccount            = addPartial("partials/settings-account.html")
	tmplPartSettingsIncomeExpenses     = addPartial("partials/settings-incomeexpenses.html")
	tmplPartSettingsDataIncomeExpenses = addPartial("partials/settings-data-incomeexpenses.html")
	tmplPartIncomeExpenseConfirmed     = addPartial("partials/incomeexpense-confirmed.html")
)

const (
	AppTitle = "Budget Helper"
	GET      = "GET"
	POST     = "POST"
	DELETE   = "DELETE"
)

func addTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path), "./templates/base.html"))
}

func addPartial(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("./templates/%v", path)))
}

func displayErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, detailedMessage string) {
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
		DetailedMessage: detailedMessage,
	}

	// TODO: Add more messages here as they are implemented in the code.
	switch statusCode {
	case http.StatusBadRequest:
		data.StatusMessage = "Bad Request"
		break
	case http.StatusMethodNotAllowed:
		data.StatusMessage = "Method Not Allowed"
		break
	case http.StatusNotFound:
		data.StatusMessage = "Not Found"
		break
	case http.StatusInternalServerError:
		data.StatusMessage = "Internal Server Error"
		break
	default:
		data.StatusMessage = "Somethign went wrong."
		break
	}

	w.WriteHeader(statusCode)
	err := tmplError.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func displayLoginRequired(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppTitle  string
		PageTitle string
	}{
		AppTitle:  AppTitle,
		PageTitle: "Login Required",
	}

	w.WriteHeader(http.StatusUnauthorized)
	err := tmplLoginRequired.ExecuteTemplate(w, "base", data)
	if err != nil {
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}
}

func (rt *Router) userIsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	session, err := rt.Store.Get(r, "session")
	if err != nil {
		log.Printf("userIsLoggedIn: %v\n", err)
		displayErrorPage(w, r, http.StatusInternalServerError,
			"Something went wrong. Please try again later.")
	}

	_, isset := session.Values["userID"]

	return isset
}

func (rt *Router) getUserIDFromSession(r *http.Request) (uint, error) {
	session, err := rt.Store.Get(r, "session")

	id, isset := session.Values["userID"]
	if !isset {
		// TODO: Generate a new error instead
		log.Fatalf("getUserIDFromSession: userID not set in session.\n")
	}

	userID, ok := id.(uint)
	if !ok {
		// TODO: Generate a new error instead
		log.Fatalf("getUserIDFromSession: Could not convert userID interface to uint.\n")
	}

	return userID, err
}

func trimmedFormValue(r *http.Request, key string) string {
	return strings.TrimSpace(r.PostFormValue(key))
}

func nameAndPassValid(username string, password string) bool {
	return strings.TrimSpace(username) != "" && strings.TrimSpace(password) != ""
}

func saveJson(data struct{}) error {
	json, err := json.MarshalIndent(data, "", "\t")
	err = os.WriteFile("./data-budget.json", json, 0666)

	return err
}

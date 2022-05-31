package router

import (
	"budget-helper/database"
	"fmt"
	"html/template"
	"net/http"
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
	tmplPartAccountConfirmed           = addPartial("partials/account-confirmed.html")
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
		return
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
		return
	}
}

func trimmedFormValue(r *http.Request, key string) string {
	return strings.TrimSpace(r.PostFormValue(key))
}

func nameAndPassValid(username string, password string) bool {
	return strings.TrimSpace(username) != "" && strings.TrimSpace(password) != ""
}

func getCurrencyString(c database.Currency) string {
	var str string

	switch c {
	case database.USD:
		str = "USD"
		break
	case database.EUR:
		str = "EUR"
		break
	case database.SEK:
		str = "SEK"
		break
	default:
		str = "Unknown currency"
		break
	}

	return str
}

func getCurrency(n uint) database.Currency {
	var currency database.Currency
	switch n {
	case 1:
		currency = database.USD
		break
	case 2:
		currency = database.EUR
		break
	case 3:
		currency = database.SEK
		break
	default:
		currency = database.UnknownCurrency
		break
	}

	return currency
}

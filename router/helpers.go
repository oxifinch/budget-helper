package router

import (
	"budget-helper/database"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

// Getting absolute path from project root to make template.Must
// compatible with the tests in this package.
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
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

// Traversing backwards from the basepath of the router/ package here, which makes it compatible with both
// running the app itself, and running tests within this package.
func addTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("%v/../templates/%v", basepath, path), fmt.Sprintf("%v/../templates/base.html", basepath)))
}

func addPartial(path string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("%v/../templates/%v", basepath, path)))
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

func budgetCategoriesAllocated(b *database.Budget) float64 {
	var bcAllocated float64

	for _, bc := range b.BudgetCategories {
		bcAllocated += bc.Allocated
	}

	return bcAllocated
}

func budgetCategoriesSpent(b *database.Budget) float64 {
	var bcSpent float64

	for _, bc := range b.BudgetCategories {
		var spentInBC float64
		for _, p := range bc.Payments {
			spentInBC += p.Amount
		}
		bcSpent += spentInBC
	}

	return bcSpent
}

func budgetBufferAllocated(b *database.Budget) float64 {
	return b.Allocated - budgetCategoriesAllocated(b)
}

func budgetBufferSpent(b *database.Budget) float64 {
	var bufSpent float64

	for _, bc := range b.BudgetCategories {
		var spentInBC float64

		for _, p := range bc.Payments {
			spentInBC += p.Amount
		}

		if spentInBC > bc.Allocated {
			bufSpent += (spentInBC - bc.Allocated)
		}
	}

	return bufSpent
}

/*
	The primary measurement for how much of a user's budget has been spent
	is calculated from the total allocated into INDIVIDUAL CATEGORIES, and
	the payments in each of those categories. The total allocated to the
	entire budget, which is divided between what is added to categories and
	whatever remains(the buffer). This is because the buffer is only meant
	as a backup, not as just another category to spend in.
*/
func budgetPercentageSpent(b *database.Budget) int {
	return int((budgetCategoriesSpent(b) / budgetCategoriesAllocated(b)) * 100)
}

func budgetCategoryPercentageSpent(bc *database.BudgetCategory) int {
	var spentInBC float64

	for _, p := range bc.Payments {
		spentInBC += p.Amount
	}

	return int((spentInBC / bc.Allocated) * 100)
}

/*
	How much has been spent of the budget's total allocated amount, which
	is what the user enters when they create their new budget, BEFORE
	allocating any money into individual categories. Since this includes
	the buffer, it's not very useful to the user as a means to get an
	accurate overview of their budget. However, it could be useful in other
	situations, so I'm including it here in case it's needed.
*/
func budgetTotalPercentageSpent(b *database.Budget) int {
	var totalSpent float64

	for _, bc := range b.BudgetCategories {
		var spentInBC float64

		for _, p := range bc.Payments {
			spentInBC += p.Amount
		}

		totalSpent += spentInBC
	}

	return int((totalSpent / b.Allocated) * 100)
}

package database

import "gorm.io/gorm"

type Currency uint

const (
	UnknownCurrency Currency = iota
	USD
	EUR
	SEK
)

type User struct {
	gorm.Model
	Username       string
	Password       string
	ActiveBudgetID uint
	Currency       Currency
	Budgets        []Budget
	Categories     []Category
	IncomeExpenses []IncomeExpense
}

type Budget struct {
	gorm.Model
	StartDate        string
	EndDate          string
	Allocated        float64
	UserID           uint
	BudgetCategories []BudgetCategory
}

type Category struct {
	gorm.Model
	Name        string
	Description string
	Color       string
	UserID      uint
	User        User
}

type BudgetCategory struct {
	gorm.Model
	Allocated  float64
	CategoryID uint
	BudgetID   uint
	Category   Category
	Payments   []Payment
}

type Payment struct {
	gorm.Model
	Date             string
	Amount           float64
	Note             string
	BudgetCategoryID uint
	BudgetCategory   BudgetCategory
}

type IncomeExpense struct {
	gorm.Model
	Label   string
	Day     uint
	Amount  float64
	Enabled bool
	UserID  uint
	User    User
}

package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string
	Password       string
	ActiveBudgetID uint
	Currency       string
	Budgets        []Budget
	Categories     []Category
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

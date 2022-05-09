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
	Allocated        float32
	UserID           uint
	BudgetCategories []BudgetCategory
}

type Category struct {
	gorm.Model
	Name        string
	Description string
	Color       string
	UserID      uint
}

type BudgetCategory struct {
	gorm.Model
	Allocated  float32
	CategoryID uint
	BudgetID   uint
	Category   Category
	Payments   []Payment
}

type Payment struct {
	gorm.Model
	Date             string
	Amount           float32
	Note             string
	BudgetCategoryID uint
	BudgetCategory   BudgetCategory
}

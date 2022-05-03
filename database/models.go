package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type Budget struct {
	gorm.Model
	StartDate string
	EndDate   string
	Allocated float32
	Currency  string
	UserID    uint
	User      User
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
	Allocated  float32
	Remaining  float32
	BudgetID   uint
	Budget     Budget
	CategoryID uint
	Category   Category
}

package budgets

import (
	"budget-helper/database"
)

type BudgetRepo struct {
	db *database.Database
}

func NewBudgetRepo(db *database.Database) *BudgetRepo {
	return &BudgetRepo{db}
}

func (b *BudgetRepo) Get(id int) (*database.Budget, error) {
	var budget database.Budget

	err := b.db.Preload("BudgetCategories.Category").
		Preload("BudgetCategories.Payments").
		Find(&budget, 1).Error

	return &budget, err
}

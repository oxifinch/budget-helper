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
		First(&budget, id).Error

	return &budget, err
}

func (b *BudgetRepo) Create(startDate string, endDate string, allocated float32) (uint, error) {

	budget := database.Budget{
		StartDate: startDate,
		EndDate:   endDate,
		Allocated: allocated,
		// Currency:  "SEK", // TODO: Currency should be set on the User instead of budget
	}

	err := b.db.Create(&budget).Error

	return budget.ID, err
}

func (b *BudgetRepo) CreateBudgetCategory(budgetID uint, categoryID uint, allocated float32) (uint, error) {

	bc := database.BudgetCategory{
		BudgetID:   budgetID,
		CategoryID: categoryID,
		Allocated:  allocated,
	}

	err := b.db.Create(&bc).Error

	return bc.ID, err
}

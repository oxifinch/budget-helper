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

func (b *BudgetRepo) Get(id uint) (*database.Budget, error) {
	var budget database.Budget

	err := b.db.Preload("BudgetCategories.Category").
		Preload("BudgetCategories.Payments").
		First(&budget, id).Error

	return &budget, err
}

func (b *BudgetRepo) GetByUserID(id uint) (*database.Budget, error) {
	var budget database.Budget

	err := b.db.
		Preload("BudgetCategories.Category").
		Preload("BudgetCategories.Payments").
		Where("user_id = ?", id).
		First(&budget).Error

	return &budget, err
}

func (b *BudgetRepo) Create(id uint, startDate string, endDate string,
	allocated float64) (uint, error) {

	budget := database.Budget{
		StartDate: startDate,
		EndDate:   endDate,
		Allocated: allocated,
		UserID:    id,
	}

	err := b.db.Create(&budget).Error

	return budget.ID, err
}

func (b *BudgetRepo) CreateBudgetCategory(budgetID uint, categoryID uint,
	allocated float64) (uint, error) {

	bc := database.BudgetCategory{
		BudgetID:   budgetID,
		CategoryID: categoryID,
		Allocated:  allocated,
	}

	err := b.db.Create(&bc).Error

	return bc.ID, err
}

// Gets only basic information about a specific budget, for when
// you don't want to load the whole thing with preloads.
func (b *BudgetRepo) GetInfo(id uint) (*database.Budget, error) {
	var budget database.Budget

	err := b.db.First(&budget, id).Error

	return &budget, err
}

package incomeexpenses

import (
	"budget-helper/database"
)

type IncomeExpenseRepo struct {
	db *database.Database
}

func NewIncomeExpenseRepo(db *database.Database) *IncomeExpenseRepo {
	return &IncomeExpenseRepo{db}
}

func (i *IncomeExpenseRepo) Get(id uint) (*database.IncomeExpense, error) {
	var incomeExpense database.IncomeExpense

	err := i.db.First(&incomeExpense, id).Error

	return &incomeExpense, err
}

func (i *IncomeExpenseRepo) GetAllWithUserID(id uint) ([]database.IncomeExpense, error) {
	var incomeExpenses []database.IncomeExpense

	err := i.db.Preload("User").
		Where("user_id = ?", id).
		Find(&incomeExpenses).Error

	return incomeExpenses, err
}

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

func (i *IncomeExpenseRepo) Create(label string, day uint, amount float64) (uint, error) {
	ie := database.IncomeExpense{
		Label:  label,
		Day:    day,
		Amount: amount,
	}

	err := i.db.Create(&ie).Error

	return ie.ID, err
}

func (i *IncomeExpenseRepo) Update(id uint, label string, day uint, amount float64) error {
	var ie database.IncomeExpense

	// Check if record exists before trying to update.
	err := i.db.Where("id = ?", id).First(&ie, id).Error
	if err != nil {

		// Return early if the record doesn't exist.
		return err
	}

	ie.Label = label
	ie.Day = day
	ie.Amount = amount
	i.db.Save(&ie)

	return err
}

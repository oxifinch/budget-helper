package payments

import (
	"budget-helper/database"
)

type PaymentRepo struct {
	db *database.Database
}

func NewPaymentRepo(db *database.Database) *PaymentRepo {
	return &PaymentRepo{db}
}

func (p *PaymentRepo) Get(id uint) (*database.Payment, error) {
	var payment database.Payment

	err := p.db.First(&payment, id).Error

	return &payment, err
}

func (p *PaymentRepo) GetAllByBudgetID(id uint) ([]database.Payment, error) {
	var payments []database.Payment

	err := p.db.Joins("BudgetCategory").
		Preload("BudgetCategory.Category").
		Where("budget_id", id).
		Order("date desc").
		Find(&payments).Error

	return payments, err
}

func (p *PaymentRepo) GetAllByBudgetCategoryID(id uint) ([]database.Payment, error) {
	var payments []database.Payment

	err := p.db.
		Preload("BudgetCategory.Category").
		Preload("BudgetCategory.Category.User").
		Where("budget_category_id = ?", id).
		Order("date desc").
		Find(&payments).Error

	return payments, err
}

func (p *PaymentRepo) GetAllByUserID(id uint) ([]database.Payment, error) {
	var payments []database.Payment

	// TODO: How do I filter out by BudgetCategory.Category.UserID = id in the query?
	// This solution doesn't work, it just select all payments regardless of UserID.
	err := p.db.
		Preload("BudgetCategory.Category").
		Preload("BudgetCategory.Category.User").
		Where(&database.BudgetCategory{
			Category: database.Category{
				UserID: id,
			},
		}).
		Order("date desc").
		Find(&payments).Error

	// Instead, this awfully inefficient solution works, for now.
	var filteredPayments []database.Payment
	for _, p := range payments {
		if p.BudgetCategory.Category.UserID == id {
			filteredPayments = append(filteredPayments, p)
		}
	}

	return filteredPayments, err
}

func (p *PaymentRepo) Create(date string, amount float64, note string, budgetCategoryID uint) (uint, error) {
	payment := database.Payment{
		Date:             date,
		Amount:           amount,
		Note:             note,
		BudgetCategoryID: budgetCategoryID,
	}

	err := p.db.Create(&payment).Error

	return payment.ID, err
}

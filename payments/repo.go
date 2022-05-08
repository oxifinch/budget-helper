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

func (p *PaymentRepo) GetAllByBudgetCategoryID(id uint) ([]*database.Payment, error) {
	var payments []*database.Payment

	err := p.db.Where("budget_category_id = ?", id).
		Find(payments).Error

	return payments, err
}

package categories

import "budget-helper/database"

type CategoryRepo struct {
	db *database.Database
}

func NewCategoryRepo(db *database.Database) *CategoryRepo {
	return &CategoryRepo{db}
}

func (c *CategoryRepo) Get(id uint) (*database.Category, error) {
	var category database.Category

	err := c.db.First(&category, id).Error

	return &category, err
}

func (c *CategoryRepo) GetAllWithUserID(id uint) ([]database.Category, error) {
	var categories []database.Category

	err := c.db.Where("user_id = ?", id).
		Find(&categories).Error

	return categories, err
}

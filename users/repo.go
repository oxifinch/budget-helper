package users

import (
	"budget-helper/database"
)

type UserRepo struct {
	db *database.Database
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Get(id uint) (*database.User, error) {
	var user database.User

	err := u.db.First(&user, id).Error

	return &user, err
}

func (u *UserRepo) GetAll() ([]database.User, error) {
	var users []database.User

	err := u.db.Find(&users).Error

	return users, err
}

func (u *UserRepo) GetByCredentials(username string, password string) (*database.User, error) {
	var user database.User

	err := u.db.Where(&database.User{
		Username: username,
		Password: password,
	}).First(&user).Error

	return &user, err
}

func (u *UserRepo) Create(username string, password string) (uint, error) {
	newUser := database.User{
		Username: username,
		Password: password,
	}

	err := u.db.Create(&newUser).Error

	return newUser.ID, err
}

func (u *UserRepo) UpdateSettings(id uint, activeBudgetID uint, currency database.Currency) error {
	var user database.User

	// Check if record exists before trying to update.
	err := u.db.Where("id = ?", id).First(&user, id).Error
	if err != nil {
		// Return early if the record doesn't exist.
		return err
	}

	user.ActiveBudgetID = activeBudgetID
	user.Currency = currency

	return u.db.Save(&user).Error
}

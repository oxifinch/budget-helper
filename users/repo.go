package users

import (
	"budget-helper/database"
	"budget-helper/models"
	"context"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepo struct {
	DB  *database.Database
	CTX context.Context
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db, context.Background()}
}

func (u *UserRepo) GetUser(id int) (*models.User, error) {
	return models.Users(Where("user_id=?", id)).One(u.CTX, u.DB)
}

func (u *UserRepo) GetAllUsers() (models.UserSlice, error) {
	return models.Users().All(u.CTX, u.DB)
}

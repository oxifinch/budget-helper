package users

import (
	"budget-helper/database"
	"budget-helper/models"
	"context"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepo struct {
	db  *database.Database
	ctx context.Context
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db, context.Background()}
}

func (u *UserRepo) GetUser(id int) (*models.User, error) {
	res, err := models.Users(Where("user_id=?", id)).One(u.ctx, u.db)

	return res, err
}

func (u *UserRepo) GetAllUsers() (models.UserSlice, error) {
	res, err := models.Users().All(u.ctx, u.db)

	return res, err
}

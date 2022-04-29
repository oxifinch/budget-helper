package users

import (
	"budget-helper/database"
	"budget-helper/models"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepo struct {
	db  *database.Database
	ctx context.Context
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db, context.Background()}
}

func (u *UserRepo) Get(id int) (*models.User, error) {
	return models.Users(Where("user_id=?", id)).One(u.ctx, u.db)
}

func (u *UserRepo) GetAll() (models.UserSlice, error) {
	return models.Users().All(u.ctx, u.db)
}

func (u *UserRepo) Create(username string, password string) (int64, error) {
	Usr := models.User{
		Username: username,
		Password: password,
	}

	// All fields have been set correctly and username and password are included.
	// It seems that boil.Infer cannot find the right columns, but using Whitelist
	// or Greylist doesn't work either. Is there something wrong with the schema?
	fmt.Printf("Usr: %v\n", Usr)
	fmt.Printf("UserId: %v\n", Usr.UserID)
	fmt.Printf("Username: %v\n", Usr.Username)
	fmt.Printf("Password: %v\n", Usr.Password)

	err := Usr.Insert(u.ctx, u.db, boil.Infer())
	return Usr.UserID, err
}

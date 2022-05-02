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

func (u *UserRepo) Get(id int) (*models.User, error) {
	return models.Users(Where("user_id=?", id)).One(u.ctx, u.db)
}

func (u *UserRepo) GetAll() (models.UserSlice, error) {
	return models.Users().All(u.ctx, u.db)
}

func (u *UserRepo) GetByCredentials(username string, password string) (*models.User, error) {
	return models.Users(Where("username = ?", username), And("password = ?", password)).One(u.ctx, u.db)
}

func (u *UserRepo) Create(username string, password string) (int64, error) {
	// TODO: Fix SQLBoiler errors! Use this method only temporarily.
	queryStr := "INSERT INTO user (username, password) VALUES (?, ?)"

	query, err := u.db.Prepare(queryStr)
	defer query.Close()

	res, err := query.Exec(username, password)

	newId, err := res.LastInsertId()

	return newId, err

	// newUser := models.User{
	// 	Username: username,
	// 	Password: password,
	// }

	// boil.DebugMode = true
	// err := newUser.Insert(u.ctx, u.db, boil.Infer())
	// return newUser.UserID, err
}

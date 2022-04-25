package users

import (
	"budget-helper/database"
	"budget-helper/models"
	"context"
	"log"
)

type UserRepo struct {
	db  *database.Database
	ctx context.Context
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db, context.Background()}
}

func (u *UserRepo) GetAllUsers() ([]*User, error) {
	res, err := models.Users().All(u.ctx, u.db)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	userList := []*User{}
	for _, val := range res {
		newUser := NewUser()

		newUser.ID = int(val.UserID)
		newUser.Username = val.Username
		newUser.Password = val.Password

		userList = append(userList, newUser)
	}

	return userList, err
}

//func (u *UserRepo) GetUser(id int) (*User, error) {
//
//	res, err := models.Users(qm.Where("user_id=?", id)).One(u.ctx, u.db)
//}

package user

import (
	"budget-helper/database"
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

func (u *UserRepo) GetAllUsers() {
	//func (u *UserRepo) GetAllUsers() (*User, error) {

	// TODO: Database is closed here. Why? Am I referencing it wrong?
	err := u.db.Ping()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// TODO: Get all users, loop through the results and print
	// to terminal for now.
	//res, err := models.Users().All(u.ctx, u.db)
	//if err != nil {
	//	log.Fatalf("error: %v\n", err)
	//}

	//for idx, val := range res {
	//	fmt.Printf("%v :: %v\n", idx, val)
	//}

}

//func (u *UserRepo) GetUser(id int) (*User, error) {
//
//	res, err := models.Users(qm.Where("user_id=?", id)).One(u.ctx, u.db)
//}

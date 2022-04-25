package users

type User struct {
	ID       int
	Username string
	Password string
}

func NewUser() *User {
	return &User{}
}

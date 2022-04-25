package users

type User struct {
	ID       int
	Username string
	Password string
}

func NewUser(id int, username string, password string) *User {
	return &User{ID: id, Username: username, Password: password}
}

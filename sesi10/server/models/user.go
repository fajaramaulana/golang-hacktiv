package models

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}

var Users = []User{}

func FindByEmail(email string) *User {
	for _, user := range Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

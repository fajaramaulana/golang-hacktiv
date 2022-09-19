package repositories

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Users = []User{
	{
		ID:        1,
		Name:      "John Doe",
		Address:   "Jl. Jalan",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func GetUser() ([]User, error) {
	if len(Users) == 0 {
		return nil, errors.New("User not found")
	}
	return Users, nil
}

func CreateUser(user User) ([]User, error) {
	newUser := append(Users, user)

	return newUser, nil
}

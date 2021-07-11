package models

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	ID        int
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("User ID must be NULL")
	}
	user.ID = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}

func GetUserById(ID int) (User, error) {
	for _, u := range users {
		if u.ID == ID {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with Id '%v' not found", ID)
}

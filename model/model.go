package model

import (
	"database/sql"
)

type User struct {
	Username string `json: "username"`
	FullName string `json: "fullname"`
	Email    string `json: "email"`
	IsActive bool   `json: "is_active"`
}

type Country struct {
	Name string `json: "name"`
}

var users = []User{
	{
		Username: "Ietry",
		FullName: "Eve Pepe",
		Email:    "ietry@example.com",
		IsActive: true,
	},
	{
		Username: "Janny",
		FullName: "Janny Pepe",
		Email:    "jenny@example.com",
		IsActive: true,
	}, {
		Username: "Pelesh",
		FullName: "Pelesh Pepe",
		Email:    "pelesh@example.com",
		IsActive: true,
	}, {
		Username: "Micah",
		FullName: "Micah Pepe",
		Email:    "micah@example.com",
		IsActive: true,
	},
}

var countries = []Country{
	{
		Name: "Nairobi",
		
	},
	{
		Name: "Kajiado",
	
	}, {
		Name: "Nakuru",

	}, {
		Name: "Turkana",
		
	},
}

func ListUsers() []User {
	return users
}

func GetUsers(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, sql.ErrNoRows
}

func CreateUser(user User) error{
	users = append(users, user)
	return nil
}

func ListCountries() []Country {
	return countries
}
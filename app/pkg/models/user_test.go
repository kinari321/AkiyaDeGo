package models

import (
	// "database/sql"
	"testing"
)

var users = []User{
	{
		Name:     "Alice",
		Email:    "alice@gmail.com",
		PassWord: "alice_pass",
	},
	{
		Name:     "Bob",
		Email:    "bob@gmail.com",
		PassWord: "bob_pass",
	},
}

func TestCreateUser(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	u, err := GetUserByEmail(users[0].Email)
	if err != nil {
		t.Errorf("User not created. err: %v", err)
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
}

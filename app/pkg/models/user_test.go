package models

import (
	"database/sql"
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

// var sessions = []Session{
// 	{
// 		user.ID
// 	}
// }

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

func TestGetUser(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	u, err := GetUserByEmail(users[0].Email)
	if err != nil {
		t.Errorf("User not created. err: %v", err)
	}
	u, err = GetUser(u.ID)
	if err != nil {
		t.Errorf("Cannot retrieve user. err:%v", err)
	}
	if u.Email != users[0].Email {
		t.Errorf("Wrong user retrieved. want:%v, got:%v", users[0], u)
	}
}

func TestUpdateUser(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	u, err := GetUserByEmail(users[0].Email)
	if err != nil {
		t.Errorf("User not created. err: %v", err)
	}
	users[0].Name = "Alice Update"
	users[0].ID = u.ID
	if err := users[0].UpdateUser(); err != nil {
		t.Errorf("Cannot update user. err:%v", err)
	}
}

func Test_UserDelete(t *testing.T) {
	cmd := `DELETE FROM users WHERE name = "Alice";`
	_, err = Db.Exec(cmd)
	cmd = `DELETE FROM users WHERE name = "Alice Update";`
	_, err = Db.Exec(cmd)

	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	u, err := GetUserByEmail(users[0].Email)
	if err != nil {
		t.Errorf("User not created. err: %v", err)
	}
	if err := u.DeleteUser(); err != nil {
		t.Errorf("Cannot delete user. err:%v", err)
	}
	_, err = GetUserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- User not deleted.")
	}
}

// ーーーーーーーーーーーーーーーーーーーーーーSESSIONーーーーーーーーーーーーーーーーーーーーーー

func TestCreateSession(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Errorf("Cannot create session. err: %v", err)
	}
	if session.UserID != users[0].ID {
		t.Error("User not linked with session")
	}
}

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
	setup()
	defer setup()
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

func TestCheckSession(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	session, err := users[0].CreateSession()
	uuid := session.UUID
	s := Session{UUID: uuid}
	valid, err := s.CheckSession()
	if err != nil {
		t.Error(err, "Cannot check session")
	}
	if valid != true {
		t.Error(err, "Session is not valid")
	}
}

func TestDeleteSessionByUUID(t *testing.T) {
	setup()
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Errorf("Cannot create session. err: %v", err)
	}
	uuid := session.UUID
	s := Session{UUID: uuid}
	err = s.DeleteSessionByUUID()
	if err != nil {
		t.Error(err, "Cannot delete session")
	}
	valid, err := s.CheckSession()
	if err == nil {
		t.Error(err, "Cannot check session")
	}
	if valid == true {
		t.Error(err, "Session is not valid")
	}
}

func GetUserBySession(t *testing.T) {
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	session, err := users[0].CreateSession()
	uuid := session.UUID
	s := Session{UUID: uuid}
	user, err := s.GetUserBySession()
	if err != nil {
		t.Error(err, "Cannot get user by session")
	}
	if users[0].UUID != user.UUID {
		t.Errorf("Wrong user by session. want:%v, got:%v", users[0].UUID, user.UUID)
	}
	defer setup()
}

func setup() {
	cmd := `truncate table users;`
	_, err = Db.Exec(cmd)
	cmd = `truncate table sessions;`
	_, err = Db.Exec(cmd)
}

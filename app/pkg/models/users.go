package models

import (
	"fmt"
	"time"

	"github.com/kinari321/AkiyaDeGo/app/errors"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Posts     []Post
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// ーーーーーーーーーーーーーーーーーーーーーーUSERーーーーーーーーーーーーーーーーーーーーーー

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
 		name,
 		email,
 		password,
 		created_at) VALUES (?, ?, ?, ?, ?)`
	uuid, err := CreateUUID()
	if err != nil {
		return errors.SetError(errors.ErrNewUUID, fmt.Sprintf("create uuid failed: %s", err))
	}
	_, err = Db.Exec(cmd,
		uuid,
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("create user failed: %s", err))
	}

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at
	FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("update user failed: %s", err))
	}

	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("delete user failed: %s", err))
	}

	return err
}

// ーーーーーーーーーーーーーーーーーーーーーーSESSIONーーーーーーーーーーーーーーーーーーーーーー

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at
	FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at) VALUES (?, ?, ?, ?)`
	uuid, err := CreateUUID()
	if err != nil {
		return session, errors.SetError(errors.ErrNewUUID, fmt.Sprintf("create uuid failed: %s", err))
	}
	_, err = Db.Exec(cmd1, uuid, u.Email, u.ID, time.Now())
	if err != nil {
		return session, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create session failed: %s", err))
	}

	cmd2 := `SELECT id, uuid, email, user_id, created_at
	FROM sessions WHERE user_id = ? AND email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT id, uuid, email, user_id, created_at
	FROM sessions WHERE uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)
	if err != nil {
		valid = false
		return valid, errors.SetError(errors.ErrDataBase, fmt.Sprintf("check session failed: %s", err))
	}
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("delete session by uuidfailed: %s", err))
	}

	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, created_at FROM users
	WHERE id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}

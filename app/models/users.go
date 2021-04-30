package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	ID        int
	Email     string
	UserID    int
	CreatedAt time.Time
}

// ーーーーーーーーーーーーーーーーーーーーーーUSERーーーーーーーーーーーーーーーーーーーーーー

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
 		name,
 		email,
 		password,
 		created_at) VALUES (?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ? email = ? WHERE id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ーーーーーーーーーーーーーーーーーーーーーーSESSIONーーーーーーーーーーーーーーーーーーーーーー

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT id, name, email, password, created_at
	FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

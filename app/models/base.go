package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

const (
	tableNameUser = "users"
	tableNamePost = "posts"
)

func init() {
	Db, err = sql.Open("mysql", "akiya:password@(127.0.0.1:3306)/akiyadego?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		name TEXT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNameUser)
	Db.Exec(cmdU)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		title TEXT NOT NULL,
		description TEXT NULL,
		user_id INT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNamePost)
	Db.Exec(cmdP)
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

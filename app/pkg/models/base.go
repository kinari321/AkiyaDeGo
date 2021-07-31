package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

var (
	Db  *sql.DB
	err error
)

const (
	tableNameUser    = "users"
	tableNamePost    = "posts"
	tableNameSession = "sessions"
)

func init() {
	//	Db, err = sql.Open("mysql", "akiya:password@tcp(127.0.0.1:3306)/akiyadego?parseTime=true") // ローカル
	Db, err = sql.Open("mysql", "akiya:password@tcp(mysql_container:3306)/akiyadego?parseTime=true") // Docker
	if err != nil {
		log.Fatalln(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		uuid TEXT NOT NULL,
		name TEXT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNameUser)
	Db.Exec(cmdU)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		imagepath TEXT NULL,
		title TEXT NOT NULL,
		type TEXT NOT NULL,
		prefecture TEXT NOT NULL,
		description TEXT NULL,
		user_id INT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNamePost)
	Db.Exec(cmdP)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		uuid TEXT NOT NULL,
		email TEXT NOT NULL,
		user_id INT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNameSession)
	Db.Exec(cmdS)

}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

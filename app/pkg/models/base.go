package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/kinari321/AkiyaDeGo/app/config"
)

var (
	Db         *sql.DB
	err        error
	USER       = config.Config.DBUser
	DBMS       = "mysql"
	DBPROTOCOL = config.Config.DBProtocol
	DBNAME     = config.Config.DBName
	DBPASS     = config.Config.DBPass
)

const (
	tableNameUser    = "users"
	tableNamePost    = "posts"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(DBMS, USER+":"+DBPASS+"@"+DBPROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	createTableUser := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		uuid TEXT NOT NULL,
		name TEXT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNameUser)
	Db.Exec(createTableUser)

	createTablePost := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		imagepath TEXT NULL,
		title TEXT NOT NULL,
		category TEXT NOT NULL,
		prefecture TEXT NOT NULL,
		description TEXT NULL,
		user_id INT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNamePost)
	Db.Exec(createTablePost)

	createTableSession := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT NOT NULL AUTO_INCREMENT,
		uuid TEXT NOT NULL,
		email TEXT NOT NULL,
		user_id INT NOT NULL,
		created_at DATETIME NULL,
		PRIMARY KEY (id));`, tableNameSession)
	Db.Exec(createTableSession)
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

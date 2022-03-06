package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/kinari321/AkiyaDeGo/app/config"
	"github.com/kinari321/AkiyaDeGo/app/errors"
)

var (
	Db         *sql.DB
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
	Db, err := sql.Open(DBMS, USER+":"+DBPASS+"@"+DBPROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Printf("database open failed: %+v\n", errors.StackTrace(err))
	}
	if err := Db.Ping(); err != nil {
		log.Printf("database connection failed: %+v\n", errors.StackTrace(err))
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
		category TEXT NOT NULL,
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

func CreateUUID() (uuid.UUID, error) {
	uuidobj, err := uuid.NewUUID()
	if err != nil {
		return uuidobj, errors.SetError(errors.ErrNewUUID, fmt.Sprintf("invalid uuid: %s", err))
	}
	return uuidobj, nil
}

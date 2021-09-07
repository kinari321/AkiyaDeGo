package config

import (
	"github.com/kinari321/AkiyaDeGo/app/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	Port       string
	DBName     string
	DBUser     string
	DBPass     string
	DBProtocol string
	LogFile    string
	Static     string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	Config = getDotEnv()
}

func getDotEnv() ConfigList {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	Config.Port = os.Getenv("PORT")
	Config.DBName = os.Getenv("DBNAME")
	Config.DBUser = os.Getenv("DBUSER")
	Config.DBPass = os.Getenv("DBPASS")
	Config.DBProtocol = os.Getenv("DBPROTOCOL")
	Config.LogFile = os.Getenv("LOGFILE")
	Config.Static = os.Getenv("STATIC")
	return Config
}

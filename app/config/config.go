package config

import (
	"github.com/kinari321/AkiyaDeGo/app/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/getsentry/sentry-go"

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
	LoadSentry()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	Config = getDotEnv()
}

func LoadSentry(){
	err := sentry.Init(sentry.ClientOptions{
		Dsn: 
	})
}

func getDotEnv() ConfigList {
	env_path := os.Getenv("ENV_PATH")
	err := godotenv.Load(env_path)
	if err != nil {
		log.Println("Error loading .env file")
	}
	Config.DBName = os.Getenv("DBNAME")
	Config.DBUser = os.Getenv("DBUSER")
	Config.DBPass = os.Getenv("DBPASS")
	Config.DBProtocol = os.Getenv("DBPROTOCOL")
	Config.LogFile = os.Getenv("LOGFILE")
	Config.Static = os.Getenv("STATIC")
	return Config
}

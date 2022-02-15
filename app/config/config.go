package config

import (
	"github.com/kinari321/AkiyaDeGo/app/utils"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

var Config ConfigList

type ConfigList struct {
	Port       string
	DBName     string
	DBUser     string
	DBPass     string
	DBProtocol string
	LogFile    string
	Static     string
	Sentry     string
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	Config = getDotEnv()
	LoadSentry()
	utils.LoggingSettings(Config.LogFile)
}

func getDotEnv() ConfigList {
	env_path := os.Getenv("ENV_PATH")
	err := godotenv.Load(env_path)
	if err != nil {
		log.Println("Error loading .env file")
	}
	//TODO: envファイルの全大文字気持ち悪い
	Config.DBName = os.Getenv("DBNAME")
	Config.DBUser = os.Getenv("DBUSER")
	Config.DBPass = os.Getenv("DBPASS")
	Config.DBProtocol = os.Getenv("DBPROTOCOL")
	Config.LogFile = os.Getenv("LOGFILE")
	Config.Static = os.Getenv("STATIC")
	Config.Sentry = os.Getenv("Sentry")

	return Config
}

func LoadSentry() {
	dsn := Config.Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer func() {
		err := recover()

		if err != nil {
			sentry.CurrentHub().Recover(err)
			sentry.Flush(time.Second * 5)
		}
	}()

}

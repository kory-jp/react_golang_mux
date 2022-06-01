package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kory-jp/react_golang_mux/api/utils"
)

type ConfigList struct {
	Env                string
	Port               string
	LogFile            string
	SQLDriver          string
	UserName           string
	Password           string
	DBHost             string
	DBPort             string
	DBName             string
	AwsBucket          string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		err := godotenv.Load("production.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
	} else {
		// p, _ := os.Getwd()
		// if string(p) == "/app/api" {
		err := godotenv.Load("env/development.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
		// } else {
		// 	err := godotenv.Load("../../../env/development.env")
		// 	if err != nil {
		// 		log.Println(err)
		// 		log.Panicln(err)
		// 	}
		// }
	}

	Config = ConfigList{
		Env:                os.Getenv("GO_ENV"),
		Port:               os.Getenv("API_PORT"),
		LogFile:            os.Getenv("LOG_FILE"),
		SQLDriver:          os.Getenv("DRIVER"),
		UserName:           os.Getenv("USER_NAME"),
		Password:           os.Getenv("PASSWORD"),
		DBHost:             os.Getenv("HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBName:             os.Getenv("DB_NAME"),
		AwsBucket:          os.Getenv("AWS_BUCKET"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsRegion:          os.Getenv("AWS_REGION"),
	}
}

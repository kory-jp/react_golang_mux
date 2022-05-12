package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kory-jp/react_golang_mux/api/utils"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	LogFile   string
	SQLDriver string
	UserName  string
	Password  string
	DBHost    string
	DBPort    string
	DBName    string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	env := os.Getenv("GO_ENV")
	if env == "production" {
		err := godotenv.Load("env/production.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
	} else {
		err := godotenv.Load("env/development.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
	}

	Config = ConfigList{
		Port:      os.Getenv("API_PORT"),
		LogFile:   cfg.Section("api").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		UserName:  os.Getenv("USER_NAME"),
		Password:  os.Getenv("PASSWORD"),
		DBHost:    os.Getenv("HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		Static:    cfg.Section("api").Key("static").String(),
	}
}

package config

import (
	"log"
	"os"

	"github.com/kory-jp/react_golang_mux/api/utils"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port       string
	LogFile    string
	SessionKey string
	SQLDriver  string
	UserName   string
	Password   string
	DBHost     string
	DBPort     string
	DBname     string
	Static     string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	var cfg *ini.File
	var err error
	env := os.Getenv("GO_ENV")
	if env == "production" {
		cfg, err = ini.Load("ini/production_config.ini")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		cfg, err = ini.Load("ini/development_config.ini")
		if err != nil {
			log.Fatalln(err)
		}
	}

	Config = ConfigList{
		Port:       cfg.Section("api").Key("port").MustString("8000"),
		LogFile:    cfg.Section("api").Key("logfile").String(),
		SessionKey: cfg.Section("api").Key("session_key").String(),
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		UserName:   cfg.Section("db").Key("user_name").String(),
		Password:   cfg.Section("db").Key("password").String(),
		DBHost:     cfg.Section("db").Key("host").String(),
		DBPort:     cfg.Section("db").Key("port").String(),
		DBname:     cfg.Section("db").Key("db_name").String(),
		Static:     cfg.Section("api").Key("static").String(),
	}
}

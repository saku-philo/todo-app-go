package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port       string
	SQLDriver  string
	DbName     string
	DbUser     string
	DbPassword string
	LogFile    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.Loggingsettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:       cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		DbName:     cfg.Section("db").Key("name").String(),
		DbUser:     cfg.Section("db").Key("user").String(),
		DbPassword: cfg.Section("db").Key("password").String(),
		LogFile:    cfg.Section("web").Key("logfile").String(),
	}
}

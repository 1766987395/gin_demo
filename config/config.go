package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	Cfg          *ini.File
	HttpRort     string
	AppMode      string
	ReadTimeOut  int
	WriteTimeOut int
	Db           string
	DbHost       string
	DBPort       string
	DbUser       string
	DbPassWord   string
	DbName       string
	DbPrefix     string
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.ini': %v", err)
	}
	loadService()
	loadMysql()
	loadApp()
}

func loadService() {
	HttpRort = Cfg.Section("service").Key("HttpRort").String()
	AppMode = Cfg.Section("service").Key("AppMode").String()
	ReadTimeOut, _ = Cfg.Section("service").Key("ReadTimeOut").Int()
	WriteTimeOut, _ = Cfg.Section("service").Key("WriteTimeOut").Int()
}

func loadMysql() {
	Db = Cfg.Section("mysql").Key("Db").String()
	DbHost = Cfg.Section("mysql").Key("DbHost").String()
	DBPort = Cfg.Section("mysql").Key("DBPort").String()
	DbUser = Cfg.Section("mysql").Key("DbUser").String()
	DbPassWord = Cfg.Section("mysql").Key("DbPassWord").String()
	DbName = Cfg.Section("mysql").Key("DbName").String()
	DbPrefix = Cfg.Section("mysql").Key("DbPrefix").String()
}

func loadApp() {
	PageSize, _ = Cfg.Section("app").Key("PageSize").Int()
	JwtSecret = Cfg.Section("app").Key("JwtSecret").String()
}

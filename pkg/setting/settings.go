package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Type	 string
	Name     string
}

type JWT struct {
	Key string
	Issuer string
	Expires int
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var JWTSetting = &JWT{}

var cfg *ini.File

func Init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("auth", JWTSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

package config

import (
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Conf Config
	Db   *gorm.DB
)

type database struct {
	Name string
	DSN  string
}

type server struct {
	ServiceName string `toml:"service_name"`
	Port        int
	Timeout     int
}

type Config struct {
	DB     database `toml:"database"`
	Server server   `toml:"server"`
}

func InitConfig(path string) {
	_, err := toml.DecodeFile(path, &Conf)
	if err != nil {
		panic(err)
	}
	Db, err = gorm.Open(mysql.Open(Conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

package config

import (
	"github.com/BurntSushi/toml"
)

const ServerConfigPathEvnKey = "user_db_rpc_config_path"

var (
	Conf Config
)

type database struct {
	Name    string `toml:"name"`
	DSN     string `toml:"dsn"`
	MaxConn int    `toml:"max_conn"`
	MaxOpen int    `toml:"max_open"`
	Timeout int64  `toml:"timeout"`
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
}

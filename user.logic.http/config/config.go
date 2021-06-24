package config

import (
	"github.com/BurntSushi/toml"
)

const ServerConfigPathEvnKey = "user_logic_http_config_path"

var (
	Conf Config
)

type server struct {
	ServiceName string `toml:"service_name"`
	Port        string
	Timeout     int
}

type Config struct {
	Server server `toml:"server"`
}

func InitConfig(path string) {
	_, err := toml.DecodeFile(path, &Conf)
	if err != nil {
		panic(err)
	}
}

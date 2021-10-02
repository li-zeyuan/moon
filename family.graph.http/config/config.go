package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const serverConfigPathEvnKey = "family_graph_http_config_path"

var (
	Conf Config
)

type server struct {
	ServiceName string `toml:"service_name"`
	Port        string `toml:"port"`
	Timeout     int    `toml:"timeout"`
}

type serverClient struct {
	ServiceName string `toml:"service_name"`
	Address     string `toml:"address"`
}

type database struct {
	Name    string `toml:"name"`
	DSN     string `toml:"dsn"`
	MaxConn int    `toml:"max_conn"`
	MaxOpen int    `toml:"max_open"`
	Timeout int64  `toml:"timeout"`
}

type Config struct {
	Server       server         `toml:"server"`
	ServerClient []serverClient `toml:"server_client"`
	DB           database       `toml:"database"`
}

func init() {
	configPath := os.Getenv(serverConfigPathEvnKey)
	_, err := toml.DecodeFile(configPath, &Conf)
	if err != nil {
		panic(err)
	}
}

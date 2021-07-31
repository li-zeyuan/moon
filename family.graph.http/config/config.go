package config

import (
	"github.com/BurntSushi/toml"
)

const ServerConfigPathEvnKey = "family_graph_http_config_path"

var (
	Conf Config
)

type server struct {
	ServiceName string `toml:"service_name"`
	Port        string
	Timeout     int
}

type serverClient struct {
	ServiceName string `toml:"service_name"`
	Address     string `toml:"address"`
}

type Config struct {
	Server       server         `toml:"server"`
	ServerClient []serverClient `toml:"server_client"`
}

func InitConfig(path string) {
	_, err := toml.DecodeFile(path, &Conf)
	if err != nil {
		panic(err)
	}
}

func GetServerClient(sClientName string) serverClient {
	for _, c := range Conf.ServerClient {
		if c.ServiceName == sClientName {
			return c
		}
	}

	return serverClient{}
}

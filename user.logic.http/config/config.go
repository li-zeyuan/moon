package config

import (
	"os"

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

type wechat struct {
	AppId  string `toml:"app_id"`
	Secret string `toml:"secret"`
}

type Config struct {
	Server       server         `toml:"server"`
	ServerClient []serverClient `toml:"server_client"`
	DB           database       `toml:"database"`
	Wechat       wechat         `toml:"wechat"`
	JwtSecretKey string         `json:"jwt_secret_key"`
}

func init() {
	configPath := os.Getenv(ServerConfigPathEvnKey)
	_, err := toml.DecodeFile(configPath, &Conf)
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

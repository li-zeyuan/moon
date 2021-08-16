package boot

import (
	"os"

	"github.com/li-zeyuan/micro/family.graph.http/config"
)

func Init() {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
	config.InitDatabase(&config.Conf)
}

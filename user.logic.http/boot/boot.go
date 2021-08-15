package boot

import (
	"os"

	"github.com/li-zeyuan/micro/user.logic.http/config"
)

func Init() {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
}

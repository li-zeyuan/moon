package boot

import (
	"os"

	"github.com/li-zeyuan/micro/user.db.rpc/config"
)

func Init() {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
	config.InitDatabase(&config.Conf)
}

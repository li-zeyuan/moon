package boot

import (
	"os"

	"github.com/li-zeyuan/micro/user.db.rpc/config"
)

func init() {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
}

package boot

import (
	"os"

	"github.com/li-zeyuan/micro/family.graph.http/config"
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
)

func Init() {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
	sequence.Init()
}

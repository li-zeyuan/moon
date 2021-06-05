package boot

import (
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
)

func Init(configPath string) {
	sequence.Init()
	config.InitConfig(configPath)
}

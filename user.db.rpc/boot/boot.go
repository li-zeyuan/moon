package boot

import (
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	_ "github.com/li-zeyuan/micro/user.db.rpc/config"
)

func Init() {
	config.InitDatabase(&config.Conf)
}

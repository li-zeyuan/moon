package boot

import (
	"github.com/li-zeyuan/micro/user.logic.http/config"
)

func Init() {
	config.InitDatabase(&config.Conf)
}

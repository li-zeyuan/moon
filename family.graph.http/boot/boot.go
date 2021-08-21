package boot

import (
	"github.com/li-zeyuan/micro/family.graph.http/config"
)

func Init() {
	config.InitDatabase(&config.Conf)
}

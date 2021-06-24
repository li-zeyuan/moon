package boot

import (
	"net/http"
	"os"

	"github.com/li-zeyuan/micro/micro.common.api/sequence"
	"github.com/li-zeyuan/micro/user.logic.http/config"
	"github.com/li-zeyuan/micro/user.logic.http/router"
)

func Init(mux *http.ServeMux) {
	configPath := os.Getenv(config.ServerConfigPathEvnKey)
	config.InitConfig(configPath)
	router.Init(mux)
	sequence.Init()
}

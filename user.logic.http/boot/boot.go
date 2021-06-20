package boot

import (
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
	"github.com/li-zeyuan/micro/user.logic.http/router"
	"net/http"
)

func Init(mux *http.ServeMux)  {
	router.Init(mux)
	sequence.Init()
}

package router

import (
	"net/http"

	"github.com/li-zeyuan/micro/user.logic.http/app/api"
)

func Init(srv *http.ServeMux) {
	srv.HandleFunc("/api/login/sing_up", api.Login.SingUp)
}

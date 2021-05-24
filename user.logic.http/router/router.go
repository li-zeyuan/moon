package router

import (
	"net/http"

	"github.con/li-zeyuan/micro/user.logic.http/app/api"
)

func Init(srv *http.ServeMux) {
	srv.HandleFunc("/hello", api.HelloHandler)
}

package router

import (
	"net/http"
	"user.logic.http/app/api"
)

func Init(srv *http.ServeMux) {
	srv.HandleFunc("/hello", api.HelloHandler)
}

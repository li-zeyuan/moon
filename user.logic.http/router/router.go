package router

import (
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/httprouter"
	middlewareCommon "github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/li-zeyuan/micro/user.logic.http/app/api"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
)

func Init(srv *http.ServeMux) {
	r := httprouter.NewRouter()
	r.Use(middlewareCommon.RequestIdMiddleware)
	r.Use(middleware.InfraMiddleware)
	r.Use(middlewareCommon.JwtMiddleware)
	// todo JwtMiddleware 可自选
	r.Add("/api/user_login/sing_up", http.HandlerFunc(api.Login.SingUp))

	for url, handler := range r.Mux {
		log.Println("api: ", url)
		srv.Handle(url, handler)
	}
}

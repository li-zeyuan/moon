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
	r.Add("/api/login/phone_code", http.HandlerFunc(api.Login.PhoneCode))
	r.Add("/api/login/phone_login", http.HandlerFunc(api.Login.PhoneLogin))
	r.Add("/api/login/wechat_login", http.HandlerFunc(api.Login.WechatLogin))

	r.Use(middlewareCommon.JwtMiddleware)
	r.Add("/api/profile/detail", http.HandlerFunc(api.Profile.Detail))
	r.Add("/api/profile/update", http.HandlerFunc(api.Profile.Update))

	for url, handler := range r.Mux {
		log.Println("api: ", url)
		srv.Handle(url, handler)
	}
}

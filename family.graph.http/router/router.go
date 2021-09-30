package router

import (
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/api"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/moon.common.api/httprouter"
	middlewareCommon "github.com/li-zeyuan/micro/moon.common.api/middleware"
)

func Init(srv *http.ServeMux) {
	r := httprouter.NewRouter()
	r.Use(middlewareCommon.RequestIdMiddleware)
	r.Use(middleware.InfraMiddleware)
	r.Add("/api/family/create", http.HandlerFunc(api.Family.Create))
	r.Add("/api/family/join", http.HandlerFunc(api.Family.Join))
	r.Add("/api/family/quit", http.HandlerFunc(api.Family.Quit))

	r.Add("/api/family_graph/create", http.HandlerFunc(api.FamilyGraph.Create))
	r.Add("/api/family_graph/detail", http.HandlerFunc(api.FamilyGraph.Detail))
	r.Add("/api/family_graph/update", http.HandlerFunc(api.FamilyGraph.Update))
	r.Add("/api/family_graph/delete", http.HandlerFunc(api.FamilyGraph.Delete))
	r.Add("/api/family_graph/graph", http.HandlerFunc(api.FamilyGraph.Graph))

	for url, handler := range r.Mux {
		log.Println("api: ", url)
		srv.Handle(url, handler)
	}
}

package router

import (
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/api"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	middleware2 "github.com/li-zeyuan/micro/micro.common.api/middleware"
)

func Init(srv *http.ServeMux) {
	r := NewRouter()
	r.Use(middleware2.RequestIdMiddleware)
	r.Use(middleware.InfraMiddleware)
	r.Add("/api/family/create", http.HandlerFunc(api.Family.Create))
	r.Add("/api/family/join", http.HandlerFunc(api.Family.Join))
	r.Add("/api/family/quit", http.HandlerFunc(api.Family.Quit))
	r.Add("/api/family_graph/create", http.HandlerFunc(api.FamilyGraph.Create))
	r.Add("/api/family_graph/update", http.HandlerFunc(api.FamilyGraph.Update))
	r.Add("/api/family_graph/detail", http.HandlerFunc(api.FamilyGraph.Detail))
	r.Add("/api/family_graph/delete", http.HandlerFunc(api.FamilyGraph.Delete))
	r.Add("/api/family_graph/graph", http.HandlerFunc(api.FamilyGraph.Graph))

	for url, handler := range r.mux {
		log.Println("api: ", url)
		srv.Handle(url, handler)
	}
}

type middlewareFunc func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middlewareFunc
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]http.Handler),
	}
}

func (r *Router) Use(m middlewareFunc) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[route] = mergedHandler
}

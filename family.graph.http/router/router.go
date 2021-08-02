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
	r.Add("/api/family_graph/graph", http.HandlerFunc(api.FamilyGraph.MethodDispatcher))

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

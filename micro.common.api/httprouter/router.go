package httprouter

import "net/http"

type middlewareFunc func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middlewareFunc
	Mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		Mux: make(map[string]http.Handler),
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

	r.Mux[route] = mergedHandler
}

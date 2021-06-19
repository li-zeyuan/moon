package middleware

import (
	"log"
	"net/http"
)

func InfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareFirst - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("MiddlerwareFirst - After Handler")

	})
}

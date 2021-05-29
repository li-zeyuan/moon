package main

import (
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/user.logic.http/router"
)

func main() {
	mux := http.NewServeMux()
	router.Init(mux)
	if err := http.ListenAndServe(":7070", mux); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

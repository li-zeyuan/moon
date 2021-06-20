package main

import (
	"github.com/li-zeyuan/micro/user.logic.http/boot"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	boot.Init(mux)
	if err := http.ListenAndServe(":7070", mux); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

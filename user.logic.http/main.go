package main

import (
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/user.logic.http/boot"
	"github.com/li-zeyuan/micro/user.logic.http/config"
)

func main() {
	mux := http.NewServeMux()
	boot.Init(mux)

	log.Printf("server: %s, port: %s", config.Conf.Server.ServiceName, config.Conf.Server.Port)
	if err := http.ListenAndServe(":"+config.Conf.Server.Port, mux); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

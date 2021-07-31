package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/li-zeyuan/micro/family.graph.http/boot"
	"github.com/li-zeyuan/micro/family.graph.http/config"
	"github.com/li-zeyuan/micro/family.graph.http/router"
	"github.com/li-zeyuan/micro/micro.common.api/logger"
)

func main() {
	boot.Init()
	srv := http.Server{
		Addr: ":" + config.Conf.Server.Port,
	}

	// 创建路由管理器
	mux := http.NewServeMux()
	router.Init(mux)
	srv.Handler = mux

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, os.Kill)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	logger.DefaultLogger.Infof("server: %s, port: %s", config.Conf.Server.ServiceName, config.Conf.Server.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.DefaultLogger.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

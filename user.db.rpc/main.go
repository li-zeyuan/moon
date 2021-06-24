package main

import (
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/user.db.rpc/app/service"
	"github.com/li-zeyuan/micro/user.db.rpc/boot"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"github.com/li-zeyuan/micro/user.db.rpc/interceptor"
	"google.golang.org/grpc"
)

func main() {
	boot.Init()
	lis, err := net.Listen("tcp", ":"+config.Conf.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.InfraUnaryServerInterceptor(),
			interceptor.RequestIDUnaryServerInterceptor(),
		)),
	)

	log.Printf("server: %s, port: %s", config.Conf.Server.ServiceName, config.Conf.Server.Port)
	profile.RegisterProfileServiceServer(s, &service.ProfileServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

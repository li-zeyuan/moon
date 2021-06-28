package main

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/li-zeyuan/micro/micro.common.api/logger"
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
		logger.DefaultLogger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.InfraUnaryServerInterceptor(),
		)),
	)

	logger.DefaultLogger.Info("server: %s, port: %s", config.Conf.Server.ServiceName, config.Conf.Server.Port)
	profile.RegisterProfileServiceServer(s, &service.ProfileServer{})
	if err := s.Serve(lis); err != nil {
		logger.DefaultLogger.Fatalf("failed to serve: %v", err)
	}
}

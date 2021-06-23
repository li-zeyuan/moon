package main

import (
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/user.db.rpc/interceptor"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/li-zeyuan/micro/user.db.rpc/app/service"
	_ "github.com/li-zeyuan/micro/user.db.rpc/boot"
	"google.golang.org/grpc"
)

const (
	port = ":7072"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.InfraUnaryServerInterceptor(),
			interceptor.RequestIDUnaryServerInterceptor(),
		)),
	)

	log.Println("port: ", port)
	profile.RegisterProfileServiceServer(s, &service.ProfileServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

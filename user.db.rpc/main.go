package main

import (
	"github.com/li-zeyuan/micro/user.db.rpc/pb/profile"
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
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			//interceptor.StreamServerInfraInterceptor(),
		)),
	)

	log.Println("port: ", port)
	profile.RegisterProfileServiceServer(s, &service.ProfileServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

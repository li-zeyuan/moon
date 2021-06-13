package main

import (
	"log"
	"net"

	"github.com/li-zeyuan/micro/micro.common.api/pb/profile"
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

	s := grpc.NewServer()
	profile.RegisterProfileServiceServer(s, &service.ProfileServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

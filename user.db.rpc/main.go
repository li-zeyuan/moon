package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"user.logic.rpc/pb"
)

const (
	port = ":7072"
)

type server struct {
	pb.UnimplementedProfileServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.UpsertReq) (*pb.UpsertResp, error) {
	log.Printf("Received: %v", in.GetUid())
	return &pb.UpsertResp{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

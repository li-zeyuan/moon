package utils

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Invoke(ctx context.Context, address, url string, in, out interface{}, opts ...grpc.CallOption) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	err = conn.Invoke(ctx, url, in, out, opts...)
	if err != nil {
		log.Fatalf("grpc invoke url: %s, error: %v", url, err)
		return err
	}

	return nil
}

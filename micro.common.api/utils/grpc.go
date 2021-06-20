package utils

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Invoke(ctx context.Context, address, url string, in, out interface{}, opts ...grpc.CallOption) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	err = conn.Invoke(ctx, url, in, out, opts...)
	if err != nil {
		log.Printf("grpc invoke url: %s, error: %v", url, err)
		return err
	}

	return nil
}

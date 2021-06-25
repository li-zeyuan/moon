package utils

import (
	"context"
	"log"

	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Invoke(ctx context.Context, address, url string, in, out interface{}, opts ...grpc.CallOption) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(RequestIDClientInterceptor()))
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Printf("did not connect: %v", err)
		return err
	}

	err = conn.Invoke(ctx, url, in, out, opts...)
	if err != nil {
		log.Printf("grpc invoke url: %s, error: %v", url, err)
		return err
	}

	return nil
}

func RequestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value(middleware.RequestId)
		if requestID, ok := value.(string); ok && requestID != "" {
			md[middleware.RequestId] = []string{requestID}
		}

		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

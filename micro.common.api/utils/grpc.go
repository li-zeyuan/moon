package utils

import (
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"google.golang.org/grpc"
)

func Invoke(baseInfra *middleware.BaseInfra, address, url string, in, out interface{}, opts ...grpc.CallOption) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(middleware.RequestIDClientInterceptor()))
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		baseInfra.Log.Errorf("did not connect: %v", err)
		return err
	}

	err = conn.Invoke(baseInfra.Context, url, in, out, opts...)
	if err != nil {
		baseInfra.Log.Errorf("grpc invoke url: %s, error: %v", url, err)
		return err
	}

	return nil
}

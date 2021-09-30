package middleware

import (
	"context"

	"github.com/li-zeyuan/micro/moon.common.api/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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

		value := ctx.Value(RequestId)
		if requestID, ok := value.(string); ok && requestID != "" {
			md[string(RequestId)] = []string{requestID}
		}

		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

func Invoke(baseInfra *BaseInfra, address, url string, in, out interface{}, opts ...grpc.CallOption) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(RequestIDClientInterceptor()))
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

func RequestIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		reqId := GetRequestIdFromIncomingContext(ctx)
		ctx = context.WithValue(ctx, RequestId, reqId)
		return handler(ctx, req)
	}
}

func GetRequestIdFromIncomingContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}

	reqId := ""
	requestIDs := md[string(RequestId)]
	if len(requestIDs) >= 1 {
		reqId = requestIDs[0]
	}

	logger.DefaultLogger.Info("get request id from incoming context: ", reqId)
	return reqId
}

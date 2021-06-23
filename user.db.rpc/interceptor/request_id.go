package interceptor

import (
	"context"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		// Set request ID for context.
		requestIDs := md[middleware.RequestId]
		if len(requestIDs) >= 1 {
			ctx = context.WithValue(ctx, middleware.RequestId, requestIDs[0])
			return handler(ctx, req)
		}

		// Generate request ID and set context if not exists.
		requestID := middleware.NewRequestId()
		ctx = context.WithValue(ctx, middleware.RequestId, requestID)
		return handler(ctx, req)
	}
}
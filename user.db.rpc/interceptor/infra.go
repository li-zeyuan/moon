package interceptor

import (
	"context"

	"github.com/li-zeyuan/micro/moon.common.api/logger"
	"github.com/li-zeyuan/micro/moon.common.api/middleware"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Infra struct {
	//Client *redis.Client
	DB *gorm.DB
	*middleware.BaseInfra
}

func GetInfra(c context.Context) *Infra {
	if c == nil {
		logger.DefaultLogger.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(middleware.InfraKey).(*Infra)
	if !ok {
		logger.DefaultLogger.Warnf("can not transfer InfraKey")
		return NewInfra(context.Background(), middleware.NewRequestId())
	}

	return infra
}

func NewInfra(ctx context.Context, requestId string) *Infra {
	return &Infra{
		config.InitDatabase(&config.Conf),
		middleware.NewBaseInfra(ctx, requestId),
	}
}

func InfraUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		reqId := middleware.GetRequestIdFromIncomingContext(ctx)
		ctx = context.WithValue(ctx, middleware.InfraKey, NewInfra(ctx, reqId))
		return handler(ctx, req)
	}
}

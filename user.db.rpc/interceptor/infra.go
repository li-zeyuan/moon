package interceptor

import (
	"context"
	"log"

	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Infra struct {
	//Client *redis.Client
	DB *gorm.DB
	//Context context.Context
	RequestId string
	Log       *logger.Logger
}

func GetInfra(c context.Context) *Infra {
	if c == nil {
		log.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(middleware.InfraKey).(*Infra)
	if !ok {
		log.Println("can not transfer InfraKey")
		return NewInfra()
	}

	return infra
}

func NewInfra() *Infra {
	infra := new(Infra)
	infra.DB = config.InitDatabase(&config.Conf)
	return infra
}

func InfraUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, middleware.InfraKey, NewInfra())
		return handler(ctx, req)
	}
}

package interceptor

import (
	"context"
	"log"

	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

type Infra struct {
	//Client *redis.Client
	DB      *gorm.DB
	//Context context.Context
	RequestId  string
}

func GetInfra(c context.Context) *Infra {
	if c == nil {
		log.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(middleware.InfraKey).(*Infra)
	if !ok {
		log.Fatal("can not transfer InfraKey")
		return NewInfra(bson.NewObjectId().Hex())
	}

	return infra
}

func NewInfra(requestID string) *Infra {
	infra := new(Infra)
	infra.RequestId = requestID
	infra.DB = config.InitDatabase(&config.Conf)
	return infra
}


func StreamServerInfraInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		ctx := stream.Context()
		infra := GetInfra(ctx)
		infra.DB = config.InitDatabase(&config.Conf)

		err = handler(srv, stream)
		return err
	}
}

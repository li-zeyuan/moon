package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"gopkg.in/mgo.v2/bson"
)

type Infra struct {
	Log *logger.Logger
	//Client *redis.Client
	//Context context.Context
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
	return &Infra{
		Log: logger.NewLogger(requestID),
	}
}

func InfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := bson.NewObjectId().Hex()
		infra := NewInfra(requestId)
		ctx := r.Context()
		ctx = context.WithValue(ctx, middleware.InfraKey, infra)
		// 设置context到r.context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

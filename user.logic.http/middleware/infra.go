package middleware

import (
	"context"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Infra struct {
	//Client *redis.Client
	//Context context.Context
	RequestId string
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
		RequestId: requestID,
	}
}

func InfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := bson.NewObjectId().Hex()
		infra := NewInfra(requestId)
		ctx := r.Context()
		ctx = context.WithValue(ctx, middleware.InfraKey, infra)

		// todo: 设置context到r.context
		//r.WithContext(ctx)
		r.Header.Set(middleware.RequestId, requestId)

		next.ServeHTTP(w, r)
		log.Println("MiddlerwareFirst - After Handler")

	})
}

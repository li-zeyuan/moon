package middleware

import (
	"context"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"gopkg.in/mgo.v2/bson"
)

type Infra struct {
	A int // 项目扩充字段
	*middleware.BaseInfra
}

func NewInfra(ctx context.Context, requestId string, a int) *Infra {
	return &Infra{
		a,
		middleware.NewBaseInfra(ctx, requestId),
	}
}

func GetInfra(c context.Context) *Infra {
	if c == nil {
		logger.DefaultLogger.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(middleware.InfraKey).(*Infra)
	if !ok {
		logger.DefaultLogger.Warnf("can not transfer InfraKey")
		return NewInfra(context.Background(), bson.NewObjectId().Hex(), 1)
	}

	return infra
}

func InfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := bson.NewObjectId().Hex()
		ctx := r.Context()
		infra := NewInfra(ctx, requestId, 1)
		ctx = context.WithValue(ctx, middleware.InfraKey, infra)
		// 设置context到r.context
		r = r.WithContext(ctx)
		w.Header().Add(middleware.InfraKey, requestId)

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"context"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"gopkg.in/mgo.v2/bson"
)

type Infra struct {
	*middleware.BaseInfra
}

func NewInfra(ctx context.Context, requestId string) *Infra {
	return &Infra{
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
		return NewInfra(context.Background(), bson.NewObjectId().Hex())
	}

	return infra
}

func InfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestId, ok := "", false
		v := ctx.Value(middleware.RequestId)
		if requestId, ok = v.(string); !ok {
			requestId = middleware.NewRequestId()
		}

		infra := NewInfra(ctx, requestId)
		ctx = context.WithValue(ctx, middleware.InfraKey, infra)
		// 设置context到r.context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

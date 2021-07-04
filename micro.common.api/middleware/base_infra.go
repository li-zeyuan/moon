package middleware

import (
	"context"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"gopkg.in/mgo.v2/bson"
)

const (
	InfraKey = "infra"
)

type BaseInfra struct {
	RequestId string // 冗余
	Log       *logger.Logger
	Context   context.Context
}

func GetBaseInfra(c context.Context, reqId string) *BaseInfra {
	if c == nil {
		logger.DefaultLogger.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(InfraKey).(*BaseInfra)
	if !ok {
		logger.DefaultLogger.Warnf("can not transfer InfraKey")
		return NewBaseInfra(context.Background(), reqId)
	}

	return infra
}

func NewBaseInfra(ctx context.Context, requestID string) *BaseInfra {
	return &BaseInfra{
		Log:       logger.NewLogger(requestID),
		RequestId: requestID,
		Context:   ctx,
	}
}

func BaseInfraMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := bson.NewObjectId().Hex()
		ctx := r.Context()
		infra := NewBaseInfra(ctx, requestId)
		ctx = context.WithValue(ctx, InfraKey, infra)
		// 设置context到r.context
		r = r.WithContext(ctx)
		w.Header().Add(InfraKey, requestId)

		next.ServeHTTP(w, r)
	})
}

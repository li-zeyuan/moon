package middleware

import (
	"context"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type requestIdCtxKey string

var RequestId requestIdCtxKey = "request_id"

func NewRequestId() string {
	return bson.NewObjectId().Hex()
}

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := NewRequestId()
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestId, requestId)

		// 设置context到r.context
		r = r.WithContext(ctx)
		w.Header().Add(string(RequestId), requestId)

		next.ServeHTTP(w, r)
	})
}

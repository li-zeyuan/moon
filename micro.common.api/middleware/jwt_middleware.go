package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/response"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

type uIdCtxKey string

var uId uIdCtxKey = "uid"

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SecretKey), nil
		})
		if err != nil {
			response.AbortWithStatusJSON(w, http.StatusForbidden, err)
			return
		}

		claims, ok := token.Claims.(*utils.JwtClaims)
		if ok && token.Valid {
			logger.DefaultLogger.Infof("uid: %d, role: %d", claims.Uid, claims.Role)
		} else {
			response.AbortWithStatusJSON(w, http.StatusForbidden, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, uId, claims.Uid)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetUid(c context.Context) int64 {
	if c == nil {
		logger.DefaultLogger.Error("content is nil")
		return 0
	}

	uid, ok := c.Value(uId).(int64)
	if !ok {
		logger.DefaultLogger.Error("can not transfer uIdCtxKey")
		return 0
	}

	return uid
}

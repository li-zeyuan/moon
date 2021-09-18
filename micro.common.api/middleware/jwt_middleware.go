package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/li-zeyuan/micro/micro.common.api/logger"
	"github.com/li-zeyuan/micro/micro.common.api/response"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

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

		if claims, ok := token.Claims.(*utils.JwtClaims); ok && token.Valid {
			logger.DefaultLogger.Infof("uid: %d, role: %d", claims.Uid, claims.Role)
		} else {
			response.AbortWithStatusJSON(w, http.StatusForbidden, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/li-zeyuan/micro/moon.common.api/logger"
	"github.com/li-zeyuan/micro/moon.common.api/response"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rc := recover(); rc != nil {
				logger.DefaultLogger.Errorf("recover panic info: %q", rc)
				debug.PrintStack()
				response.AbortWithStatusJSON(w, http.StatusOK, errors.New("recover panic"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"
	"testing"
)

func TestRecoverMiddleware(t *testing.T) {
	var myHandlerFunc = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		panic("my panic")
	})
	RecoverMiddleware(myHandlerFunc)
	myHandlerFunc(nil, nil)
}

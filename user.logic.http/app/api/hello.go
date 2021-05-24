package api

import (
	"io"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	err := utils.Invoke(r.Context())
	if err != nil {
		return
	}

	io.WriteString(w, "Hello, world!\n")
}

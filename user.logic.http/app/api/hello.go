package api

import (
	"io"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/model/rpc"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	profileRpcReq := rpc.UpsertReq{}
	profileRpcResp := rpc.UpsertResp{}
	err := utils.Invoke(r.Context(), rpc.AddressProfileServer, rpc.UrlProfileUpsert, &profileRpcReq, &profileRpcResp)
	if err != nil {
		return
	}

	io.WriteString(w, "Hello, world!\n")
}

package api

import (
	"github.com/li-zeyuan/micro/micro.common.api/pb/profile"
	"io"
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/model/rpc"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

var Profile = new(profileAPI)

type profileAPI struct{}

func (p *profileAPI) Create(w http.ResponseWriter, r *http.Request) {
	profileRpcReq := profile.UpsertReq{}
	profileRpcResp := profile.UpsertResp{}
	err := utils.Invoke(r.Context(), rpc.AddressProfileServer, rpc.UrlProfileUpsert, &profileRpcReq, &profileRpcResp)
	if err != nil {
		return
	}

	io.WriteString(w, "Hello, world!\n")
}

package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/model/rpc"
	"github.com/li-zeyuan/micro/micro.common.api/pb/profile"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/library"
)

var Login = new(loginAPI)

type loginAPI struct{}

func (l *loginAPI) SingUp(w http.ResponseWriter, r *http.Request) {
	apiReq := new(model.LoginApiSingUpReq)
	err := library.ParseBody(r, apiReq)
	if err != nil {
		library.AbortWithStatusJSON(w, http.StatusInternalServerError, err)
		return
	}

	profileRpcReq := profile.UpsertReq{}
	profileRpcResp := profile.UpsertResp{}
	err = utils.Invoke(r.Context(), rpc.AddressProfileServer, rpc.UrlProfileUpsert, &profileRpcReq, &profileRpcResp)
	if err != nil {
		return
	}

	library.Json(w, http.StatusOK, "ok")
}

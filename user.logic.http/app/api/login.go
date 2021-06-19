package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/service"
	"github.com/li-zeyuan/micro/user.logic.http/library/request"
	"github.com/li-zeyuan/micro/user.logic.http/library/response"
)

var Login = new(loginAPI)

type loginAPI struct{}

func (l *loginAPI) SingUp(w http.ResponseWriter, r *http.Request) {
	apiReq := new(model.LoginApiSingUpReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = service.Login.VerifySingUp(apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusInternalServerError, err)
		return
	}

	//profileRpcReq := profile.UpsertReq{}
	//profileRpcResp := profile.UpsertResp{}
	//err = utils.Invoke(r.Context(), rpc.AddressProfileServer, rpc.UrlProfileUpsert, &profileRpcReq, &profileRpcResp)
	//if err != nil {
	//	return
	//}

	response.Json(w, http.StatusOK, "ok")
}

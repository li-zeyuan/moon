package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/service"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
	"github.com/li-zeyuan/micro/user.logic.http/library/request"
	"github.com/li-zeyuan/micro/user.logic.http/library/response"
)

var Login = new(loginAPI)

type loginAPI struct{}

func (l *loginAPI) SingUp(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())
	infra.Log.Infof("lizeyuan")

	apiReq := new(model.LoginApiSingUpReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Login.VerifySingUp(apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Login.SingUp(r.Context(), apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, nil)
}

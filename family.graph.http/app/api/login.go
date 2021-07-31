package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/service"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/family.graph.http/library/request"
	"github.com/li-zeyuan/micro/family.graph.http/library/response"
)

var Login = new(loginAPI)

type loginAPI struct{}

func (l *loginAPI) SingUp(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.LoginApiSingUpReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Login.VerifySingUp(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Login.SingUp(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, nil)
}

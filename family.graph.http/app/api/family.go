package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/service"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/request"
	"github.com/li-zeyuan/micro/micro.common.api/response"
)

var Family = new(familyAPI)

type familyAPI struct{}

func (l *familyAPI) Create(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyAPICreateReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Family.Create(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, nil)
}

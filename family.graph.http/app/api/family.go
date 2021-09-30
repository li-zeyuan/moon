package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/service"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/moon.common.api/request"
	"github.com/li-zeyuan/micro/moon.common.api/response"
)

var Family = new(familyAPI)

type familyAPI struct{}

// Create
// @Summary 创建家族
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyAPICreateReq true " "
// @Router /api/family/create [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
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

	response.Json(w, http.StatusOK, struct{}{})
}

// Join
// @Summary 加入家族
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyAPIJoinReq true " "
// @Router /api/family/join [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *familyAPI) Join(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyAPIJoinReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Family.Join(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

// Quit
// @Summary 退出家族
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyAPIQuitReq true " "
// @Router /api/family/quit [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *familyAPI) Quit(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyAPIQuitReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.Family.Quit(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/service"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/request"
	"github.com/li-zeyuan/micro/micro.common.api/response"
)

var FamilyGraph = new(familyGraphAPI)

type familyGraphAPI struct{}

// Create
// @Summary 创建族谱图节点
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyGraphAPICreateReq true " "
// @Router /api/family_graph/create [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *familyGraphAPI) Create(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyGraphAPICreateReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.FamilyGraph.CreateNode(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

// Detail
// @Summary 族谱图节点详情
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyGraphAPIDetailReq true " "
// @Router /api/family_graph/detail [post]
// @Success 200 {object} model.FamilyGraphAPIDetailResp
func (l *familyGraphAPI) Detail(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyGraphAPIDetailReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	detail, err := service.FamilyGraph.DetailNode(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, detail)
}

// Update
// @Summary 更新族谱图节点
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyGraphAPIUpdateReq true " "
// @Router /api/family_graph/update [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *familyGraphAPI) Update(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyGraphAPIUpdateReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.FamilyGraph.UpdateNode(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

// Delete
// @Summary 删除族谱图节点
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyGraphAPIDelReq true " "
// @Router /api/family_graph/delete [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *familyGraphAPI) Delete(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyGraphAPIDelReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	err = service.FamilyGraph.DelNode(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

// Graph
// @Summary 族谱图
// @tags 家族模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.FamilyGraphAPIGraphReq true " "
// @Router /api/family_graph/graph [post]
// @Success 200 {object} model.FamilyGraphAPIGraphResp
func (l *familyGraphAPI) Graph(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.FamilyGraphAPIGraphReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	resp, err := service.FamilyGraph.GetGraph(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, resp)
}

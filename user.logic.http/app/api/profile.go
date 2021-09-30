package api

import (
	"net/http"

	commonMiddleware "github.com/li-zeyuan/micro/moon.common.api/middleware"
	"github.com/li-zeyuan/micro/moon.common.api/request"
	"github.com/li-zeyuan/micro/moon.common.api/response"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/service"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
)

var Profile = new(profileAPI)

type profileAPI struct{}

// Detail
// @Summary 获取个人信息
// @tags 用户资料模块
// @Description
// @Accept  json
// @Produce  json
// @Router /api/profile/detail [post]
// @Success 200 {object} model.ProfileApiDetailResp
func (l *profileAPI) Detail(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())
	uid := commonMiddleware.GetUid(r.Context())

	resp, err := service.Profile.Detail(infra, uid)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, resp)
}

// Update
// @Summary 更新个人信息
// @tags 用户资料模块
// @Description
// @Accept  json
// @Produce  json
// @Router /api/profile/update [post]
// @Param req body model.ProfileApiUpdateRep true " "
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *profileAPI) Update(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())
	uid := commonMiddleware.GetUid(r.Context())

	apiReq := new(model.ProfileApiUpdateRep)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}
	err = service.Profile.Update(infra, uid, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	response.Json(w, http.StatusOK, struct{}{})
}

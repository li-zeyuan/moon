package api

import (
	"net/http"

	"github.com/li-zeyuan/micro/micro.common.api/request"
	"github.com/li-zeyuan/micro/micro.common.api/response"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/service"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
)

var Login = new(loginAPI)

type loginAPI struct{}

func (l *loginAPI) PhoneCode(w http.ResponseWriter, r *http.Request) {

}

// PhoneLogin
// @Summary 手机号登录
// @tags 登录模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.LoginApiPhoneLoginReq true " "
// @Router /api/login/phone_login [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *loginAPI) PhoneLogin(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.LoginApiPhoneLoginReq)
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

// WechatLogin
// @Summary 微信登录
// @tags 登录模块
// @Description
// @Accept  json
// @Produce  json
// @Param req body model.LoginApiWeChatLoginReq true " "
// @Router /api/user_login/sing_up [post]
// @Success 200 {object} string "{"dm_error":0,"error_msg":"","data":{}}"
func (l *loginAPI) WechatLogin(w http.ResponseWriter, r *http.Request) {
	infra := middleware.GetInfra(r.Context())

	apiReq := new(model.LoginApiWeChatLoginReq)
	err := request.ParseBody(r, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	token, err := service.Login.WeChatLogin(infra, apiReq)
	if err != nil {
		response.AbortWithStatusJSON(w, http.StatusOK, err)
		return
	}

	resp := new(model.LoginApiWeChatLoginResp)
	resp.Token = token
	response.Json(w, http.StatusOK, resp)
}

package response

import (
	"encoding/json"
	"net/http"

	"github.com/li-zeyuan/micro/moon.common.api/errorenum"
)

type JsonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func AbortWithStatusJSON(w http.ResponseWriter, code int, err error) {
	if err == nil {
		return
	}

	resp := JsonResponse{}
	resp.Code = -1
	resp.Msg = err.Error()
	if errEnum, ok := err.(errorenum.ErrorCode); ok {
		resp.Code = errEnum.Code
		resp.Msg = errEnum.Msg
	}

	writeResponse(w, code, resp)
}

func Json(w http.ResponseWriter, code int, data interface{}) {
	resp := JsonResponse{}
	resp.Data = data

	writeResponse(w, code, resp)
}

func writeResponse(w http.ResponseWriter, code int, resp interface{}) {
	body, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

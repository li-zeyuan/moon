package library

import (
	"encoding/json"
	"net/http"
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
	body, _ := json.Marshal(resp)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func Json(w http.ResponseWriter, code int, data interface{}) {
	resp := JsonResponse{}
	resp.Data = data
	body, _ := json.Marshal(resp)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

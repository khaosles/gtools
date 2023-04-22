package gresponse

import (
	"fmt"
)

/*
   @File: ErrorCode.go
   @Author: khaosles
   @Time: 2023/3/7 21:54
   @Desc:
*/

type JsonResponse struct {
	// code
	Code int `json:"code" default:"0"`
	// response information
	Msg string `json:"msg" default:""`
	// data
	Data interface{} `json:"data,omitempty" default:"nil"`
	// whether success
	Success bool `json:"success" default:"false"`
}

func (r JsonResponse) setCode(code int) JsonResponse {
	r.Code = code
	return r
}

func (r JsonResponse) setMsg(Msg string) JsonResponse {
	r.Msg = Msg
	return r
}

func (r JsonResponse) setSuccess(success bool) JsonResponse {
	r.Success = success
	return r
}

func (r JsonResponse) setData(data interface{}) JsonResponse {
	r.Data = data
	return r
}

// Yes is run successful
func (r JsonResponse) Yes(data interface{}) JsonResponse {
	return r.setCode(SUCCESS.Code).setMsg(SUCCESS.Msg).setSuccess(true).setData(data)
}

// No is run failed
func (r JsonResponse) No(err Status) JsonResponse {
	return r.setCode(err.Code).setMsg(err.Msg)
}

// Custom result
func (r JsonResponse) Custom(code int, Msg string, data interface{}, success bool) JsonResponse {
	return r.setCode(code).setMsg(Msg).setData(data).setSuccess(success)
}

// CatchErr 异常捕获
func (r JsonResponse) CatchErr(err interface{}) JsonResponse {
	return r.No(CUSTOM_ERROR.SetMsg(fmt.Sprintf("%v", err)))
}

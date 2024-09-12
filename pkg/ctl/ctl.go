package ctl

import (
	"demo/pkg/e"
)

// 基础resp
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// 返回数组形式
type Datalist struct {
	Item  interface{} `json:"items"`
	Total int64       `json:"total"`
}

// 登录时使用，返回携带用户信息和token的responese

// 携带追踪信息的resp
type TrackedErrorResponse struct {
	Response
	TraceId string `json:"traceId"`
}

func RespSuccess(code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   "Ok",
		Msg:    e.GetMsg(status),
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}

func RespList(items interface{}, total int64) *Response {
	return &Response{
		Status: e.Success,
		Data: Datalist{
			Item:  items,
			Total: total,
		},
		Msg: e.GetMsg(e.Success),
	}
}

func RespError(err error, data interface{}, code ...int) *TrackedErrorResponse {
	status := e.Error
	if code != nil {
		status = code[0]
	}
	return &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Data:   data,
			Msg:    e.GetMsg(status),
			Error:  err.Error(),
		},
		TraceId: "", //TODO
	}
}

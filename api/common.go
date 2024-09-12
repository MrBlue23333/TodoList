package api

import (
	"demo/pkg/ctl"
	"demo/pkg/e"
	"encoding/json"
)

func ErrorResponse(err error) *ctl.TrackedErrorResponse {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "JSON type mismatched")
	}
	return ctl.RespError(err, "Invalid params", e.InvalidParams)
}

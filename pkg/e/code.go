package e

import "net/http"

const (
	Success       = http.StatusOK
	Error         = http.StatusInternalServerError
	InvalidParams = http.StatusBadRequest
)

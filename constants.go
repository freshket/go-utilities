package utilities

import (
	"net/http"
)

const (
	OK                           = "SUCCESS"
	DOMAIN_EXCEPTION             = "DOMAIN_EXCEPTION"
	NOT_FOUND_EXCEPTION          = "NOT_FOUND"
	UNEXPECTED_EXCEPTION         = "UNEXPECTED_EXCEPTION"
	REQUEST_VALIDATION_EXCEPTION = "REQUEST_VALIDATION_EXCEPTION"
	UNAUTHORIZED_REQUEST         = "UNAUTHORIZED"
	REQUIRES_CREDENTIALS         = "MISSING_CREDENTIALS"
	INVALID_CREDENTIALS          = "INVALID_CREDENTIALS"
	INVALID_ACCESS               = "INVALID_ACCESS"
	CONCURRENCY_EXCEPTION        = "CONCURRENCY_EXCEPTION"
	DOWNSTREAM_EXCEPTION         = "DOWNSTREAM_EXCEPTION"
)

var HttpCodeMaps = map[string]int{
	OK:                           http.StatusOK,
	DOMAIN_EXCEPTION:             http.StatusUnprocessableEntity,
	NOT_FOUND_EXCEPTION:          http.StatusNotFound,
	UNEXPECTED_EXCEPTION:         http.StatusInternalServerError,
	REQUEST_VALIDATION_EXCEPTION: http.StatusBadRequest,
	UNAUTHORIZED_REQUEST:         http.StatusUnauthorized,
	REQUIRES_CREDENTIALS:         http.StatusUnauthorized,
	INVALID_CREDENTIALS:          http.StatusUnauthorized,
	INVALID_ACCESS:               http.StatusForbidden,
	CONCURRENCY_EXCEPTION:        http.StatusConflict,
	DOWNSTREAM_EXCEPTION:         http.StatusServiceUnavailable,
}

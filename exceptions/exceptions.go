package exceptions

import (
	"fmt"
	"net/http"

	"github.com/freshket/go-utilities/constants"
)

type ApplicationError struct {
	Code    string
	Message string
}

func (ae *ApplicationError) Error() string {
	return fmt.Sprintf("error captured [%s] => %s", ae.Code, ae.Message)
}

var (
	NotFoundException           = ApplicationError{Code: constants.NOT_FOUND_EXCEPTION, Message: "requested resource does not exist"}
	UnhandledException          = ApplicationError{Code: constants.UNEXPECTED_EXCEPTION, Message: "unhandled exception"}
	DownstreamException         = ApplicationError{Code: constants.DOWNSTREAM_EXCEPTION, Message: "downstream system error"}
	RequestValidationException  = ApplicationError{Code: constants.REQUEST_VALIDATION_EXCEPTION, Message: "invalid request"}
	UnauthorisedException       = ApplicationError{Code: constants.UNAUTHORIZED_REQUEST, Message: "unauthorised"}
	InvalidAccessException      = ApplicationError{Code: constants.INVALID_ACCESS, Message: "invalid access"}
	ConcurrencyException        = ApplicationError{Code: constants.CONCURRENCY_EXCEPTION, Message: "concurrency exception"}
	MissingCredentialsException = ApplicationError{Code: constants.REQUIRES_CREDENTIALS, Message: "required credentials"}
	InvalidCredentialException  = ApplicationError{Code: constants.INVALID_CREDENTIALS, Message: "invalid credentials"}
	UnsupportedContentException = ApplicationError{Code: constants.UNSUPPORTED_CONTENT, Message: "unsupported content"}
)

var HttpStatusMap = map[string]int{
	NotFoundException.Code:           http.StatusNotFound,
	UnhandledException.Code:          http.StatusInternalServerError,
	DownstreamException.Code:         http.StatusServiceUnavailable,
	RequestValidationException.Code:  http.StatusBadRequest,
	UnauthorisedException.Code:       http.StatusUnauthorized,
	InvalidAccessException.Code:      http.StatusForbidden,
	ConcurrencyException.Code:        http.StatusConflict,
	MissingCredentialsException.Code: http.StatusUnauthorized,
	InvalidCredentialException.Code:  http.StatusUnauthorized,
	UnsupportedContentException.Code: http.StatusUnsupportedMediaType,
}

func GetHttpStatusForCode(code string) int {
	return HttpStatusMap[code]
}

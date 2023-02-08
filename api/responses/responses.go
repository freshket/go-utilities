package responses

import "github.com/freshket/go-utilities/constants"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status string      `json:"status"`
	Error  ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Stack   string `json:"stack,omitempty"`
}

func Ok(message string, payload interface{}) *Response {
	return &Response{
		Status:  constants.OK,
		Message: message,
		Data:    payload,
	}
}

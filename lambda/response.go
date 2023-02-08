package lambda

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/freshket/go-utilities/constants"
	"github.com/freshket/go-utilities/exceptions"
)

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

func CreateResponse(statusCode int, payload interface{}) events.APIGatewayProxyResponse {
	json, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(json),
	}
}

func Ok(message string, payload interface{}) events.APIGatewayProxyResponse {
	return CreateResponse(http.StatusOK, Response{
		Status:  constants.OK,
		Message: message,
		Data:    payload,
	})
}

func OkWithCode(statusCode int, message string, payload interface{}) events.APIGatewayProxyResponse {
	return CreateResponse(statusCode, Response{
		Status:  "SUCCESS",
		Message: message,
		Data:    payload,
	})
}

func Fail(err error) events.APIGatewayProxyResponse {
	var appErr *exceptions.ApplicationError
	if errors.As(err, &appErr) {
		return CreateResponse(exceptions.GetHttpStatusForCode(appErr.Code), ErrorResponse{
			Status: appErr.Code,
			Error: ErrorDetail{
				Message: appErr.Message,
				Stack:   appErr.Error(),
			},
		})
	}

	return CreateResponse(http.StatusInternalServerError, ErrorResponse{
		Status: constants.UNEXPECTED_EXCEPTION,
		Error: ErrorDetail{
			Message: constants.DEFAULT_ERROR_MESSAGE,
			Stack:   err.Error(),
		},
	})
}

package responses

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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
		Status:  OK,
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
	var appErr *ApplicationError
	if errors.As(err, &appErr) {
		return CreateResponse(GetHttpStatusForCode(appErr.Code), ErrorResponse{
			Status: appErr.Code,
			Error: ErrorDetail{
				Message: appErr.Message,
				Stack:   appErr.Error(),
			},
		})
	}

	return CreateResponse(http.StatusInternalServerError, ErrorResponse{
		Status: UNEXPECTED_EXCEPTION,
		Error: ErrorDetail{
			Message: "unhandled exception",
			Stack:   err.Error(),
		},
	})
}

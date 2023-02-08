package middlewares

import (
	"errors"
	"net/http"

	"github.com/freshket/go-utilities/api/responses"
	"github.com/freshket/go-utilities/constants"
	"github.com/freshket/go-utilities/exceptions"
	"github.com/labstack/echo"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var appErr *exceptions.ApplicationError

	if errors.As(err, &appErr) {
		c.JSON(exceptions.GetHttpStatusForCode(appErr.Code), responses.ErrorResponse{
			Status: appErr.Code,
			Error: responses.ErrorDetail{
				Message: appErr.Message,
				Stack:   appErr.Error(),
			},
		})

		return
	}

	c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
		Status: constants.UNEXPECTED_EXCEPTION,
		Error: responses.ErrorDetail{
			Message: constants.DEFAULT_ERROR_MESSAGE,
			Stack:   err.Error(),
		},
	})
}

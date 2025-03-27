package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata"`
}

func SuccessResponse(c echo.Context, message string, data interface{}, metadata interface{}) error {

	return c.JSON(http.StatusOK, &APIResponse{
		Status:   "success",
		Message:  message,
		Data:     data,
		Metadata: metadata,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {

	return c.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
	})
}

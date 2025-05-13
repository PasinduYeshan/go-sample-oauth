package merchantservices

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/common/response"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {

	e.GET("", GetAllServices)
}

func GetAllServices(c echo.Context) error {

	return response.SuccessResponse(c, "Ads retrieved successfully", SampleData, nil)
}

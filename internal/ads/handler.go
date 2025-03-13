package ads

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/ads", GetAllAds)
}

func GetAllAds(c echo.Context) error {

	return c.JSON(http.StatusOK, ads)
}

// Sample data.
var ads = []Ad{
	{ID: 1, Title: "Ad 1", Description: "Description for Ad 1", Price: 100.00},
	{ID: 2, Title: "Ad 2", Description: "Description for Ad 2", Price: 200.00},
	{ID: 3, Title: "Ad 3", Description: "Description for Ad 3", Price: 300.00},
}

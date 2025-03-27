package ads

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/common/response"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {

	e.GET("", GetAllAds)
}

func GetAllAds(c echo.Context) error {

	return response.SuccessResponse(c, "Ads retrieved successfully", ads, nil)
}

// Sample data.
var ads = []Ad{
	{ID: 1, Title: "Ad 1", Description: "Description for Ad 1", Price: 100.00},
	{ID: 2, Title: "Ad 2", Description: "Description for Ad 2", Price: 200.00},
	{ID: 3, Title: "Ad 3", Description: "Description for Ad 3", Price: 300.00},
}

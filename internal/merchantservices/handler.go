package merchantservices

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {

	e.GET("", GetAllServices)
}

func GetAllServices(c echo.Context) error {

	return c.JSON(http.StatusOK, services)
}

// Sample data.
var services = []Service{
	{
		ID:          1,
		Name:        "Car Oil Change",
		Description: "Complete engine oil change and filter replacement",
		Price:       59.99,
		Category:    "Car Repair",
	},
	{
		ID:          2,
		Name:        "Bike Tune-up",
		Description: "Full bicycle tune-up including gear and brake adjustment",
		Price:       45.00,
		Category:    "Bike Repair",
	},
	{
		ID:          3,
		Name:        "Car Brake Repair",
		Description: "Brake pad replacement and system check",
		Price:       129.99,
		Category:    "Car Repair",
	},
}

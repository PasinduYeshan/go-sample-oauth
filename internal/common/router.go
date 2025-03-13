package common

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/ads"
	"github.com/PasinduYeshan/go-sample-oauth/internal/merchantservices"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes Registers all routes for the application.
func RegisterRoutes(e *echo.Echo) {

	apiV1 := e.Group("/api/v1")
	adsGroup := apiV1.Group("/ads")
	merchantServicesGroup := apiV1.Group("/services")

	ads.RegisterRoutes(adsGroup)
	merchantservices.RegisterRoutes(merchantServicesGroup)
}

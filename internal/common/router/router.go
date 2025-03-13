package router

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/ads"
	"github.com/PasinduYeshan/go-sample-oauth/internal/merchantservices"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes Registers all routes for the application.
func RegisterRoutes(e *echo.Echo) {

	ads.RegisterRoutes(e)
	merchantservices.RegisterRoutes(e)
}

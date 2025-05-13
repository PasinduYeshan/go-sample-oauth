package main

import (
	"log"

	"github.com/PasinduYeshan/go-sample-oauth/internal/common/router"
	"github.com/PasinduYeshan/go-sample-oauth/internal/config" // Added import for config package

	"github.com/labstack/echo/v4"
)

func main() {
	// Load API scope configuration
	_, err := config.LoadScopeConfig("config/api_scopes.yaml") // Path relative to project root
	if err != nil {
		log.Fatalf("Could not load API scope configuration: %v", err)
	}
	log.Println("API scope configuration loaded successfully.")

	log.Println("Server is running on port 8080")
	if err := BuildServer().Start(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func BuildServer() *echo.Echo {

	e := echo.New()
	router.RegisterRoutes(e)
	return e
}

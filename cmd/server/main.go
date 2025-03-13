package main

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/common"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	log.Println("Server is running on port 8080")
	if err := BuildServer().Start(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func BuildServer() *echo.Echo {

	e := echo.New()
	common.RegisterRoutes(e)
	return e
}

package main

import (
	"os"

	"github.com/brunoksato/go-api-lambda/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", api.Home)

	private := e.Group("/api")
	private.GET("/test", api.Test)

	// Start server
	addr := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(addr))
}

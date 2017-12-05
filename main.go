package main

import (
	"os"

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
	e.GET("/", Home)

	// Start server
	addr := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(addr))
}

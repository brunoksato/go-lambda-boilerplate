package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	hello := map[string]interface{}{"status": "Hi API"}
	return c.JSON(http.StatusOK, hello)
}

func Test(c echo.Context) error {
	hello := map[string]interface{}{"status": "Test API"}
	return c.JSON(http.StatusOK, hello)
}

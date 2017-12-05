package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	hello := map[string]interface{}{"status": "GRUFF API"}
	return c.JSON(http.StatusOK, hello)
}

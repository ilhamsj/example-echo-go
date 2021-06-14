package routes

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func init() *echo.Echo  {
	
	e := echo.New()

	e.GET("/", func (c echo.Context) error  {
		return c.String(http.StatusOK, "Hello, Echo !")
	})

	return e
}
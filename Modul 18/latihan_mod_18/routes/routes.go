package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	fmt.Println("===== Init Route ===== ")

	e := echo.New()

	// Hello World
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e = InitRouteBuku(e)

	return e
}

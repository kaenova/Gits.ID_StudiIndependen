package routes

import (
	"latihan_mod_18/controllers"

	"github.com/labstack/echo/v4"
)

func InitRouteBuku(e *echo.Echo) *echo.Echo {
	e.GET("/buku", controllers.GetBuku)
	e.POST("/buku", controllers.InsertBuku)
	return e
}

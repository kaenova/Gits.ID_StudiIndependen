package controllers

import (
	"fmt"
	"latihan_mod_18/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetBuku(c echo.Context) error {
	var (
		res Response = NewResponse()
	)

	data, err := models.GetAllBuku()
	if err != nil {
		fmt.Println(err.Error())
		res.Status = http.StatusInternalServerError
		return c.JSON(res.Status, res)
	}

	res.Data = data
	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}

func InsertBuku(c echo.Context) error {
	var (
		res Response = NewResponse()
	)

	nama := c.FormValue("nama")
	if strings.TrimSpace(nama) == "" {
		return c.JSON(res.Status, res)
	}

	data, err := models.CreateBuku(nama)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(res.Status, res)
	}

	res.Data = data
	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}

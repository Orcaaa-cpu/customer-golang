package controllers

import (
	"net/http"
	"strconv"

	"github.com/Orcaaa-cpu/customer-golang/models"
	"github.com/labstack/echo"
)

func GetCustomer(c echo.Context) error {
	result, err := models.GetCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func PostCustomer(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	tanggal := c.FormValue("tanggal")

	result, err := models.PostCustomer(nama, alamat, telepon, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCustomer(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	tanggal := c.FormValue("tanggal")

	conId, _ := strconv.Atoi(id)

	result, err := models.UpdateCustomer(conId, nama, alamat, telepon, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCustomer(c echo.Context) error {
	id := c.FormValue("id")

	conId, _ := strconv.Atoi(id)

	result, err := models.DeleteCustomer(conId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

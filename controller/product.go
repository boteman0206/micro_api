package controller

import (
	"github.com/labstack/echo/v4"
	"micro_api/models"
	"net/http"
)

func GetProduct(c echo.Context) error {

	var res = models.Response{}

	return c.JSON(http.StatusOK, res)

}

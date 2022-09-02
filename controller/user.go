package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"micro_api/config"
	"micro_api/models"
	"net/http"
)

func GetDcUserUrl() string {
	var url string
	if client, ok := config.ClientInfo[config.DcUser]; ok {
		url = client.Addr + ":" + client.Port
	}
	return url
}

func GetUser(c echo.Context) error {

	var res = models.Response{}

	url := GetDcUserUrl()
	fmt.Println("getUser的ip和端口：", url)

	return c.JSON(http.StatusOK, res)
}

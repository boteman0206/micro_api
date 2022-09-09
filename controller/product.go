package controller

import (
	"github.com/labstack/echo/v4"
	"micro_api/config"
	"micro_api/micro_proto/pc"
	"micro_api/models"
	"net/http"
)

func GetDcProductUrl() string {
	var url string
	if client, ok := config.ClientInfo[config.DcProduct]; ok {
		url = client.Addr + ":" + client.Port
	}
	return url
}

func GetProduct(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
	}

	url := GetDcProductUrl()
	//调用product的rpc方法
	client := pc.GetDcProductGrpcClient(url)
	dto := pc.GetProductDto{
		Id:   12,
		Name: "",
		Sort: "",
	}
	product, err := client.RPC.GetProduct(client.Ctx, &dto)
	if err != nil {
		res.Msg = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Data = product

	return c.JSON(http.StatusOK, res)

}

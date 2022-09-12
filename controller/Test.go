package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"micro_api/micro_common/utils"
	"micro_api/micro_proto/pc"
	"micro_api/models"
	"net/http"
)

func Test01(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "i am fine",
	}

	log.Println("log：", utils.JsonToString(res))
	utils.BufferLog.Println("NewLog：", utils.JsonToString(res))
	fmt.Println("buf：", utils.LogBuf.String())
	return c.JSON(http.StatusOK, res)
}

func Test02(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "test02",
	}

	// 测试日志 这里打印的路径是common包的需要获取当前路径才行
	utils.MyLog.Debug("debug: ", " test02： ", utils.JsonToString(res))
	utils.MyLog.Info("Info: ", " test02： ", utils.JsonToString(res))
	utils.MyLog.Error("Error: ", " test02： ", utils.JsonToString(res))

	return c.JSON(http.StatusOK, res)
}

func Test03(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "Test03",
	}

	// 测试日志 这里打印的路径是common包的需要获取当前路径才行
	url := GetDcProductUrl()
	//调用product的rpc方法
	client := pc.GetDcProductGrpcClient(url)
	dto := pc.GetProductDto{
		Id:   12,
		Name: "",
		Sort: "",
	}
	product, err := client.RPC.TestProduct(client.Ctx, &dto)
	if err != nil {
		res.Msg = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	utils.MyLog.Info("Info: ", " test03： ", utils.JsonToString(product.Data))
	res.Data = product.Data
	return c.JSON(http.StatusOK, res)
}

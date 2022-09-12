package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"micro_api/micro_common/es"
	"micro_api/micro_common/middle"
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
	es.BufferLog.Println("NewLog：", utils.JsonToString(res))
	fmt.Println("buf：", es.LogBuf.String())

	traceId := c.Get(middle.TraceId)
	fmt.Println("traceId: ", traceId)

	return c.JSON(http.StatusOK, res)
}

func Test02(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "test02",
	}

	// 测试日志 这里打印的路径是common包的需要获取当前路径才行
	es.MyLog.Debug("debug: ", " test02： ", utils.JsonToString(res))
	es.MyLog.Info("Info: ", " test02： ", utils.JsonToString(res))
	es.MyLog.Error("Error: ", " test02： ", utils.JsonToString(res))

	return c.JSON(http.StatusOK, res)
}

func Test03(c echo.Context) error {
	// 从echo中获取trace_id
	es.MyLog.Echo = c

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "Test03",
	}

	// 测试日志 这里打印的路径是common包的需要获取当前路径才行
	url := GetDcProductUrl()
	//调用product的rpc方法
	client := pc.GetDcProductGrpcClient(url, c)
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
	es.MyLog.Info("Info: ", " test03： ", utils.JsonToString(product.Data))
	res.Data = product.Data
	return c.JSON(http.StatusOK, res)
}

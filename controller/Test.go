package controller

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"micro_api/models"
	"micro_api/utils"
	"net/http"
)

/**
golang自带的log包使用：默认是打印时间+日志信息： 可以调整
*/

var NewLog *log.Logger
var buf bytes.Buffer

func init() {
	// 自定义log输出到buffer也可以定义到文件
	NewLog = log.New(&buf, "[info: ]", log.LstdFlags)
}

func Test01(c echo.Context) error {

	var res = models.Response{
		Msg:   "success",
		Total: 0,
		Data:  "i am fine",
	}

	log.Println("log：", utils.JsonToString(res))
	NewLog.Println("NewLog：", utils.JsonToString(res))
	fmt.Println("buf：", buf.String())
	return c.JSON(http.StatusOK, res)
}

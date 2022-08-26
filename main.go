package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"micro_api/config"
	"micro_api/routes"
	"micro_api/services"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func init() {
	// 设置时区
	sh, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = sh
}

func main() {

	// 初始化配置文件
	config.InitConfig()

	go func() {
		//pprof性能分析
		profPort := fmt.Sprintf(":%d", config.ConfigRes.Ser.PprofPort)
		err := http.ListenAndServe(profPort, nil)
		if err != nil {
			return
		}
	}()

	flag.Parse()

	//初始化mysql和redis连接
	services.SetupDB()
	//关闭mysql和redis连接
	defer services.CloseDB()

	e := echo.New()
	e.Use(middleware.BodyLimit("5M"))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//错误中间件
	e.Use(middleware.Recover())

	r := e.Group("/api")
	r.GET("/test", func(c echo.Context) error {

		return c.String(http.StatusOK, "my test ....")
	})

	routes.AuthRoute(r)

	port := fmt.Sprintf(":%d", config.ConfigRes.Ser.HttpPort)

	e.Logger.Fatal(e.Start(port))

}

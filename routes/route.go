package routes

import (
	"github.com/labstack/echo/v4"
	"micro_api/controller"
)

//路由
func AuthRoute(e *echo.Group) *echo.Group {

	testGroup(e)

	//用户中心
	userGroup(e)

	//商品中心
	productGroup(e)
	return e
}

func testGroup(e *echo.Group) {

	g := e.Group("/test")
	g.GET("/myTest", controller.Test01)
}

func userGroup(e *echo.Group) {
	g := e.Group("/user")
	g.GET("/getUser", controller.GetUser)
}

func productGroup(e *echo.Group) {
	g := e.Group("/product")
	g.GET("/getProduct", controller.GetProduct)

}

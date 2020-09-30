package route

import (
	"github.com/labstack/echo/v4"
	"lee_wechat/controller"
)

//路由
func Route(e *echo.Echo) *echo.Echo {
	//登录
	e.GET("/ping", controller.Query)

	return e
}

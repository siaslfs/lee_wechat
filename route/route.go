package route

import (
	"github.com/labstack/echo/v4"
	"lee_wechat/controller"
)

//路由
func Route(e *echo.Echo) *echo.Echo {
	//登录
	e.GET("/ping", controller.Query)
	GroupFromWechat(e)
	return e
}
func GroupFromWechat(e *echo.Echo) {
	w := e.Group("wechat")
	w.GET("/receive/check", controller.CheckSignature)  //校验服务器合法性
	w.POST("/receive/check", controller.CheckSignature) //校验服务器合法性
	w.POST("/receive/msg", controller.ReceiveMsgFromWechat)
}

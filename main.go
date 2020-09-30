package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "lee_wechat/docs"
	"lee_wechat/route"
	"time"
)

func init() {
	sh, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = sh
}

// @title 项目接口文档
// @version 1.0
// @description 微信相关接口文档
// @host localhost:11001
func main() {
	//日志命令行参数化处理，可以启用禁用控制台日志等，defer确认
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//错误中间件
	e.Use(middleware.Recover())

	e = route.Route(e)

	addr := ":11001" //默认启动端口
	e.Logger.Fatal(e.Start(addr))
}

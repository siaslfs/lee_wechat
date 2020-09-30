package controller

import (
	"github.com/labstack/echo/v4"
)

// @Summary test
// @Tags test
// @Accept plain
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string	"fail"
// @Router /demo/ping [get]
func Query(c echo.Context) error {
	return c.JSON(200, "ok")
}

package api

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := "test"
		return c.JSON(fasthttp.StatusOK, v)
	}
}

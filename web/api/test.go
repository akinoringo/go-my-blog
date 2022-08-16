package api

import (
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		name := claims.Name
		return c.String(fasthttp.StatusOK, "Welcome "+name+"!")
	}
}

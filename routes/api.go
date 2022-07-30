package routes

import (
	"go-my-blog/web/api"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")

	{
		g.GET("/test", api.Test())
	}
}

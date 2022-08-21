package routes

import (
	"go-my-blog/handler/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")

	{
		g.POST("/signup", api.SignUp())
		g.POST("/auth/login", api.Login())
		g.GET("/auth/me", api.Me())
	}

	r := g.Group("/restricted")
	r.Use(middleware.JWTWithConfig(api.Config))
	r.GET("", api.Restricted())
	r.GET("/article/:id", api.GetArticle())
	r.POST("/article/create", api.Create())
}

package api

import (
	"go-my-blog/middlewares"
	"go-my-blog/models"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		article := new(models.Article)
		if err := c.Bind(article); err != nil {
			return err
		}

		uid := GetUserIDFromToken(c)
		user := new(models.User)

		if dbs.DB.Table("users").Where(models.User{ID: uid}).First(&user).RecordNotFound() {
			return echo.ErrNotFound
		}

		article.UserID = uid
		dbs.DB.Create(&article)

		return c.JSON(fasthttp.StatusOK, article)
	}
}

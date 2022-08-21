package api

import (
	"go-my-blog/middlewares"
	"go-my-blog/models"
	"strconv"

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

func GetArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		article := new(models.Article)
		if dbs.DB.Table("articles").Where(models.Article{ID: articleID}).First(&article).RecordNotFound() {
			return echo.ErrNotFound
		}

		return c.JSON(fasthttp.StatusOK, article)
	}
}

package api

import (
	"go-my-blog/middlewares"
	"go-my-blog/models"
	"log"
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

func UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		uid := GetUserIDFromToken(c)
		user := new(models.User)
		if dbs.DB.Table("users").Where(models.User{ID: uid}).First(&user).RecordNotFound() {
			return echo.ErrNotFound
		}

		updateArticle := new(models.Article)
		if err := c.Bind(updateArticle); err != nil {
			return err
		}

		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		article := new(models.Article)
		result := dbs.DB.Model(&article).Where("id = ?", articleID).Updates(models.Article{Title: updateArticle.Title, Content: updateArticle.Content})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		dbs.DB.Where("id = ?", articleID).Take(&article)
		return c.JSON(fasthttp.StatusOK, article)
	}
}

package api

import (
	"go-my-blog/models"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		tests := []models.TestType{}
		test1 := models.TestType{ID: "1", Name: "ジョナサン・ジョースター"}
		test2 := models.TestType{ID: "2", Name: "ディオ・ブランドー"}
		tests = append(tests, test1)
		tests = append(tests, test2)

		return c.JSON(fasthttp.StatusOK, tests)
	}
}

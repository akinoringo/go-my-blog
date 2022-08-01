package main

import (
	"fmt"
	"net/http"

	"go-my-blog/databases"
	"go-my-blog/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/test", testHandler)
	_, err := databases.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("database connected")
	e.Logger.Fatal(e.Start(":8080"))
}

// echoの初期設定
// インスタンスの生成
// 各種middlewareの設定
func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	routes.Init(e)

	return e
}

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

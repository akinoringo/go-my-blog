package api

import (
	"crypto/sha1"
	"fmt"
	"go-my-blog/middlewares"
	"go-my-blog/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valyala/fasthttp"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("SECRET_KEY")

var Config = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: signingKey,
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		r := new(models.User)
		if err := c.Bind(r); err != nil {
			log.Println(err)
		}

		user := models.User{}

		if dbs.DB.Table("users").Where(models.User{Email: r.Email}).First(&user).RecordNotFound() {
			user = models.User{Name: r.Name, Email: r.Email, Password: Encrypt(r.Password)}
			dbs.DB.Create(&user)
		}

		user.Password = ""

		return c.JSON(fasthttp.StatusOK, user)
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		r := new(models.User)
		if err := c.Bind(r); err != nil {
			log.Println(err)
		}
		user := models.User{}

		dbs.DB.Table("users").Where(models.User{Email: r.Email}).First(&user)

		if user.ID == 0 || user.Password != Encrypt(r.Password) {
			return &echo.HTTPError{
				Code:    fasthttp.StatusUnauthorized,
				Message: "invalid name or password",
			}
		}

		claims := &JwtCustomClaims{
			user.Name,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(signingKey)
		if err != nil {
			return err
		}

		return c.JSON(fasthttp.StatusOK, echo.Map{
			"token": tokenString,
		})
	}
}

func Me() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(fasthttp.StatusOK, echo.Map{
			"user": "user",
		})
	}
}

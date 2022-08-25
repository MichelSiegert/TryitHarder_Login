package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//required to avoid coarse issues
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowMethods},
	}))
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
	}

	e.POST("/insert", func(c echo.Context) error {

		user := parseUser(c)
		return createNewAccount(user, db, c)
	})
	e.POST("/login", func(c echo.Context) error {
		var user userData
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		return tryLogin(user, db, c)
	})

	//setPort
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

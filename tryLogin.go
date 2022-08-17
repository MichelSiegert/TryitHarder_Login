package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	//base example for GET
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	//BASE Example for POST
	e.POST("/nice", func(c echo.Context) error {
		var users []userData
		read := c.FormValue("name")
		response := &Response{Mail: "jsomichel", httpstatus: 200, Message: read, Data: "all systems ready!", User: users}

		db := connectDB()
		createTable(db)
		query, err := db.Query("SELECT *  FROM Users")
		if err != nil {
			fmt.Println(err)
			return c.JSON(response.httpstatus, response)
		}

		var a, b, d, e, f, g, h string

		defer func(query *sql.Rows) {
			err := query.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(query)
		for query.Next() {
			err := query.Scan(&a, &b, &h, &d, &e, &f, &g)
			if err != nil {
				fmt.Println(err)
			}
			tmp := userData{a, b, h, d, e, f, g}
			response.User = append(response.User, tmp)

		}
		err = query.Err()
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(response.httpstatus, response)
	})

	//set port

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

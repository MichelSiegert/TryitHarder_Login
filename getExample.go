package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func get(c echo.Context) error {
	return c.HTML(http.StatusOK, "Hello, Docker! <3")
}

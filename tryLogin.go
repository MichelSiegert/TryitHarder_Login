package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func tryLogin(user userData, db *sql.DB, c echo.Context) error {
	var response Response
	data, token := checkUser(user, db)
	if token == "-1" {
		response.Message = token
		response.Data = data
		response.httpstatus = http.StatusOK
		response.Mail = user.Email
	}

	return c.JSON(response.httpstatus, response)
}

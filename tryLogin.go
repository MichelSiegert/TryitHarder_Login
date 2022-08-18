package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func tryLogin(user userData, db *sql.DB, c echo.Context) error {
	var response Response
	token := checkUser(user, db)
	if token == "-1" {
		response.Data += "failed to login! \n"
		response.httpstatus = http.StatusPaymentRequired
	} else {
		response.Message = token
		response.Mail = user.Email
	}
	return c.JSON(response.httpstatus, response)
}

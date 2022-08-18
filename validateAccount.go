package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func createNewAccount(data userData, db *sql.DB, c echo.Context) error {
	pw := data.Password
	response := Response{httpstatus: http.StatusOK}

	if len(pw) < 7 && !strings.ContainsAny(pw, "1234567890") && !strings.Contains(pw, "abcdeghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		response.httpstatus = http.StatusBadRequest
		response.Message = "password doesnt live up to our standards!"
		return c.JSON(response.httpstatus, response)
	}

	response.User = append(response.User, data)
	response.Data = "12345 some othe stuff!"
	response.Message = "success!"
	response.Mail = "ka lol."
	insertUser(db, data)
	return c.JSON(response.httpstatus, response)
}

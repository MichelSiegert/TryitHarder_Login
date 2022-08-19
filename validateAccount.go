package main

import (
	"database/sql"
	"net/http"
	"strings"
)

func createNewAccount(data userData, db *sql.DB, c echo.Context) error {
	pw := data.Password
	response := Response{httpstatus: http.StatusOK}

	if len(pw) < 7 && !strings.ContainsAny(pw, "1234567890") && !strings.Contains(pw, "abcdeghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		response.httpstatus = http.StatusBadRequest
		response.Message = "password doesnt live up to our standards!"
		response.User = append(response.User, data)
		response.Mail = data.Email
		response.Data = "unable to create new account: Password is too short, or does not align with the the security standards of our website!"
		return c.JSON(response.httpstatus, response)
	}

	response.User = append(response.User, data)
	response.Message = "Hello " + data.Name
	response.Mail = data.Email
	response = insertUser(db, data, response)
	return c.JSON(response.httpstatus, response)
}

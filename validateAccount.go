package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func createNewAccount(data userData, db *sql.DB, c echo.Context) error {
	pw := data.Password
	if len(pw) < 7 && !strings.ContainsAny(pw, "1234567890") && !strings.Contains(pw, "abcdeghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return c.JSON(http.StatusBadRequest, struct{ status string }{status: "password didnt live up to standards: at least 7 character and 1 at least one number and letter"})
	}

	insertUser(db, data)
	return c.JSON(http.StatusOK, struct{ status string }{status: "A new User has been created!"})
}

package main

import "github.com/labstack/echo/v4"

func parseUser(ctx echo.Context) userData {
	var data userData
	data.Address = ctx.FormValue("address")
	data.Name = ctx.FormValue("name")
	data.Password = ctx.FormValue("password")
	data.Email = ctx.FormValue("email")
	data.SessID = generateRandomString()

	return data
}

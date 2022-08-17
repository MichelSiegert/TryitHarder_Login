package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	var users []userData
	read := c.FormValue("name")
	response := &Response{Mail: "jsomichel", httpstatus: 200, Message: read, Data: "all systems ready!", User: users}

	db := connectDB()
	status := 0

	newUser := userData{Name: "michel", Email: "msiegert@dons.usfca.edu", DateOfBirth: "12.03.1997", Password: "nicepassword", Address: "Am Deepenbrook 1"}
	status = insertUser(db, newUser)
	if status != 200 {
		fmt.Println("something went wrong while creating the user!")
	}
	var user userData
	user = selectUser(db, "msiegert@dons.usfca.edu")

	fmt.Println(user.Email, user.Id, user.Address)
	status = UpdateUser(db, user.Id, user.Name, user.Password, "brodersdorfer straße", user.Email, user.DateOfBirth)
	if status != 200 {
		fmt.Println("something went wrong while editing the user data")
	}
	status = DeleteUser(db, user.Id)
	if status != 200 {
		fmt.Println("something went wrong while deleting the user!")
	}
	response.User = append(users, user)
	return c.JSON(response.httpstatus, response)
}

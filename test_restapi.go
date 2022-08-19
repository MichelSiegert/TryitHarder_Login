package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
)

func testRest(c echo.Context, db *sql.DB) error {
	var users []userData
	read := c.FormValue("name")
	response := Response{Mail: "jsomichel", httpstatus: 200, Message: read, Data: "all systems ready!", User: users}

	status := 200

	newUser := userData{Name: "michel", Email: "msiegert@dons.usfca.edu", Password: "nicepassword", Address: "Am Deepenbrook 1"}
	insertUser(db, newUser, response)

	var user userData
	user = selectUser(db, "msiegert@dons.usfca.edu")

	fmt.Println(user.Email, user.Id, user.Address)
	status = UpdateUser(db, user.Id, user.Name, user.Password, user.Email, "brodersdorfer straße", user.SessID)
	if status != 200 {
		fmt.Println("something went wrong while editing the user data")
	}
	user = selectUser(db, "msiegert@dons.usfca.edu")
	fmt.Println(user.Address)
	status = DeleteUser(db, user.Id)
	if status != 200 {
		fmt.Println("something went wrong while deleting the user!")
	}
	response.User = append(users, user)
	return c.JSON(response.httpstatus, response)
}

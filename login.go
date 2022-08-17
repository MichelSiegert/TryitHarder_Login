package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
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

	var id, name, mail, address, dateOfBirth, firstLogin, password string

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(query)
	for query.Next() {
		err := query.Scan(&id, &name, &password, &mail, &address, &dateOfBirth, &firstLogin)
		if err != nil {
			fmt.Println(err)
		}
		tmp := userData{id, name, password, mail, address, dateOfBirth, firstLogin}
		response.User = append(response.User, tmp)

	}
	err = query.Err()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(response.httpstatus, response)
}

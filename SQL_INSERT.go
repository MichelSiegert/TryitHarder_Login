package main

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user userData, response Response) Response {

	query, err := db.Query("INSERT INTO USERS (username, usrpassword, email, address, session_token) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.Address, user.SessID)

	if err != nil {
		response.Data += err.Error() + "\n"
		fmt.Println(err)
		return response
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			response.Data += err.Error() + "\n"
		}
	}(query)
	if err != nil {
		fmt.Println(err)
		response.Data += err.Error() + "\n"
		return response
	}
	fmt.Println("successfully added the user!")
	response.Data += "successfully added the user!\n"
	return response
}

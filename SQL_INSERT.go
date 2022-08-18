package main

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user userData, response Response) int {

	query, err := db.Query("INSERT INTO USERS (username, usrpassword, email, address, session_token) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.Address, user.SessID)

	if err != nil {
		response.Data += err.Error() + "\n"
		fmt.Println(err)
		return 500
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
		return 500
	}
	fmt.Println("successfully added the user!")
	response.Data += "successfully added the user!\n"
	return 200
}

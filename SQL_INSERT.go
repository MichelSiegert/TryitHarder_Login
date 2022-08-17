package main

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user userData) int {

	query, err := db.Query("INSERT INTO USERS (username, usrpassword, email, address, DateOfBirth ) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.Address, user.DateOfBirth)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return 500
	}

	return 200
}

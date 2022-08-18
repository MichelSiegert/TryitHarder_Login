package main

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user userData) int {

	query, err := db.Query("INSERT INTO USERS (username, usrpassword, email, address, sessID ) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.Address, user.SessID)

	if err != nil {
		fmt.Println(err)
		return 500
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
		}
	}(query)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	fmt.Println("successfully added the user!")
	return 200
}

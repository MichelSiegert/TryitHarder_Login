package main

import (
	"database/sql"
	"fmt"
)

func UpdateUser(db *sql.DB, id string, username string, password string, email string, address string, sessID string) int {

	query, err := db.Query("UPDATE USERS SET username = ? , usrpassword = ?, email = ?, address = ?, sessID = ? WHERE id = ?", username, password, email, address, sessID, id)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(query)
	if err != nil {
		fmt.Println(err)
		return 500
	}

	return 200
}

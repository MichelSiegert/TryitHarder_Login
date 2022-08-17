package main

import (
	"database/sql"
	"fmt"
)

func DeleteUser(db *sql.DB, id string) int {

	query, err := db.Query("DELETE FROM USERS WHERE ID = ?", id)
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

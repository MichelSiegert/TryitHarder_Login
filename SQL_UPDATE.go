package main

import (
	"database/sql"
	"fmt"
)

func UpdateUser(db *sql.DB, id string, field string, newVal string) int {

	query, err := db.Query("UPDATE USERS SET ? = ?  WHERE ID = ?", field, newVal, id)
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

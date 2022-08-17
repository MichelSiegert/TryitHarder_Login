package main

import (
	"database/sql"
	"fmt"
)

func selectUser(db *sql.DB, email string) *userData {
	var id, name, mail, address, dateOfBirth, firstLogin, password string
	user := userData{}
	query, err := db.Query("SELECT * FROM USERS WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		return &user
	}
	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return &user
	}
	for query.Next() {
		err := query.Scan(&id, &name, &mail, &address, &dateOfBirth, &firstLogin, &password)
		if err != nil {
			fmt.Println(err)
			return &user
		}
	}
	err = query.Err()
	if err != nil {
		fmt.Println(err)
		return &user
	}
	return &user
}

package main

import (
	"database/sql"
	"fmt"
)

func createTable(db *sql.DB) {
	query, err := db.Query(`
	CREATE TABLE IF NOT EXISTS USERS (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		username VARCHAR(255) DEFAULT NULL,
		usrpassword VARCHAR(255) NOT NULL,
		email VARCHAR(255) DEFAULT NULL,
		address VARCHAR(255) DEFAULT NULL,
		DateOfBirth VARCHAR(255) DEFAULT '01.01.1970',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		CONSTRAINT UQ_USER_email UNIQUE (email, username)
	);`)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var id, value string
	for query.Next() {
		err := query.Scan(&id, &value)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s: %s\n", id, value)
	}
	err = query.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

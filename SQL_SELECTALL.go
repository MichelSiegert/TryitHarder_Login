package main

import (
	"database/sql"
	"fmt"
)

func selectAll(db *sql.DB) []userData {
	var data []userData
	query, err := db.Query("SELECT * FROM Users")
	if err != nil {
		fmt.Println(err)
		return data
	}

	var id, name, mail, address, dateOfBirth, firstLogin, password string

	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return data
	}

	for query.Next() {
		err := query.Scan(&id, &name, &password, &mail, &address, &dateOfBirth, &firstLogin)
		if err != nil {
			fmt.Println(err)
			return data
		}
		tmp := userData{id, name, password, mail, address, dateOfBirth, firstLogin}
		data = append(data, tmp)
		fmt.Println("a")
	}

	err = query.Err()
	if err != nil {
		fmt.Println(err)
		return data
	}
	fmt.Println(data[0].Name)
	return data
}

package main

import (
	"database/sql"
	"fmt"
)

func checkUser(usr userData, db *sql.DB) string {
	var id, name, mail, address, password, sessID string
	user := userData{}
	query, err := db.Query("SELECT * FROM USERS WHERE email = ? AND usrpassword = ? ", usr.Email, usr.Password)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(query)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	for query.Next() {
		err := query.Scan(
			&id,
			&name,
			&password,
			&mail,
			&address,
			&sessID)
		user.Name = name
		user.Id = id
		user.Email = mail
		user.Address = address
		user.SessID = sessID
		user.Password = password
		fmt.Println(name, id, mail, address, sessID, password)
		if err != nil {
			fmt.Println(err)
			return "-1"
		}

	}
	err = query.Err()
	if err != nil {
		fmt.Println(err)
		return "-1"
	}

	if user.Password == usr.Password && user.Email == usr.Email {
		fmt.Println("Done!")
		return user.SessID
	}
	return "-1"
}

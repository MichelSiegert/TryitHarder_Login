package main

import (
	"database/sql"
	"fmt"
)

func selectUser(db *sql.DB, email string) userData {
	var id, name, mail, address, dateOfBirth, firstLogin, password string
	user := userData{}
	fmt.Print("#")
	query, err := db.Query("SELECT * FROM USERS WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Print("#")
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(query)
	fmt.Print("#")
	if err != nil {
		fmt.Println(err)
		return user
	}
	for query.Next() {
		err := query.Scan(&id, &name, &mail, &address, &dateOfBirth, &firstLogin, &password)
		user.Name = name
		user.Id = id
		user.Email = mail
		user.Address = address
		user.DateOfLogin = firstLogin
		user.DateOfBirth = dateOfBirth
		user.Password = password
		fmt.Println(name, id, mail, address, firstLogin, dateOfBirth, password)
		if err != nil {
			fmt.Println(err)
			return user
		}

	}
	fmt.Print("#")
	err = query.Err()
	fmt.Print("#")
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println("# Done!")
	return user
}
